package main

import (
    "github.com/xddbom/cache-api/db"
    "github.com/xddbom/cache-api/routes"

    "github.com/gin-gonic/gin"
)

func main() {    
    var rdb = db.RedisInit()
    r := gin.Default()

    r.Use(routes.RedisMiddleware(rdb))

    routes.RoutesSetup(r, rdb)

    r.Run(":8080")
}
