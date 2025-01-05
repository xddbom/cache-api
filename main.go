package main

import (
	"github.com/xddbom/cache-api/db"
    "github.com/gin-gonic/gin"
)

func main() {    
    var rdb = db.RedisInit()

    r := gin.Default()

    r.GET("/", func(c *gin.Context){
        c.JSON(200, gin.H{
            "message": "Welcome to Cache API!",
        })
    })

    r.GET("/health", func(c *gin.Context){
        db.HealthCheck(c, rdb)
    })

    r.Run(":8080")
}