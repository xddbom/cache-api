package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func RedisInit() {
	client := redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "",
        DB:		  0,  
        Protocol: 2, 
    })

	Ctx := context.Background()

	err := client.Ping(Ctx).Err()
	if err != nil {
		panic(fmt.Sprintf("Ошибка подключения к Redis: %v", err))
	}
	
	err = client.Set(Ctx, "Connection", "successful!", 0).Err()
    if err!= nil {
        panic(err)
    }

	val, err := client.Get(Ctx, "Connection").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("Connection", val)
}