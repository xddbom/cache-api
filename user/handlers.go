package user

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)


// POST
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

	const UserTTL = 60 * time.Second
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
	err = rdb.Set(ctx, redisKey, userData, UserTTL).Err()
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


// GET
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

	const UserTTL = 60 * time.Second
	if err := rdb.Expire(ctx, redisKey, UserTTL).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update TTL for user",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}


// DELETE
func DeleteUser(c* gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rdb, ok := c.MustGet("rdb").(*redis.Client)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect to Redis",
		})
		return
	}

	userID := c.Param("id")
	redisKey := "user:" + userID

	deleted, err := rdb.Del(ctx, redisKey).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user from Redis",
		})
		return
	}

	if deleted == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
		"user_id": userID,
	})
}