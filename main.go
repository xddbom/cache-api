package main

import (
	"github.com/xddbom/cache-api/db"
    "github.com/gin-gonic/gin"
)

func main() {    
    db.RedisInit()

    r := gin.Default()

    r.GET("/", func(c *gin.Context){
        c.JSON(200, gin.H{
            "message": "Welcome to Cache API!",
        })
    })

    r.Run(":8080")
}