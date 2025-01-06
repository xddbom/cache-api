package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/xddbom/cache-api/db"
	"github.com/xddbom/cache-api/user"
)

func RoutesSetup(r *gin.Engine, rdb *redis.Client) {
	r.GET("/", func(c *gin.Context){
        c.JSON(200, gin.H{
            "message": "Welcome to Cache API!",
        })
    })

    r.GET("/health", func(c *gin.Context){
        db.HealthCheck(c, rdb)
    })

    r.POST("/users", user.CreateUser)
}