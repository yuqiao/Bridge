package main

import (
	"fmt"
	"menteslibres.net/gosexy/redis"
)

func main() {
	client := redis.New()
	err := client.Connect("127.0.0.1", 6379)

	fmt.Printf("err: %s\n", err)

	data, err := client.Get("name")

	if err == nil {
		fmt.Println("data:" + data)
	}
}
