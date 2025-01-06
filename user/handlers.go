package user

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)


func CreateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rdb, ok := c.MustGet("rdb").(*redis.Client)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect to Redis",
		})
		return
	}


	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	userData, err := json.Marshal(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to serialize user data",
		})
		return
	}

	redisKey := "user:" + user.ID
	err = rdb.Set(ctx, redisKey, userData, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save user in Redis",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user_id": user.ID,
	})
}


func GetUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rdb, ok := c.MustGet("rdb").(*redis.Client)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect to Redis",
		})
	return
	}

	userID := c.Param("id") // ?
	redisKey := "user:" + userID

	userData, err := rdb.Get(ctx, redisKey).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve user from Redis",
		})
		return
	}

	var user User
	if err := json.Unmarshal([]byte(userData), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to deserialize user data",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}