package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "192.168.0.86:6379",
		DB:   0,
	})
	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)

	// 设置键，带过期时间
	err = redisClient.Set("username", "charles", time.Hour).Err()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("设置字符 username 成功")
	}
	// 设置键，不带过期时间
	err = redisClient.Set("user", "charles", 0).Err()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("设置字符 user 成功")
	}

	// 通过键查询值
	value, err := redisClient.Get("user").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("获取字符 user 成功：", value)
	}

	// 设置集合
	err = redisClient.SAdd("test", "123456").Err()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("设置集合 test 成功")
	}

	// fmt.Printf("set :%s", set)
	flag := redisClient.SetNX("test1", "123456", time.Hour)
	fmt.Printf("flag :%s", flag)
}
