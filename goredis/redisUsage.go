package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:8080",
		Password: "",
		DB:       0,
	})

	err := client.Set("hello", "hellovalue", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("hello").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("hello", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
