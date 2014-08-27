package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func main() {
	urlStr := "mq://wangxin:sddf@10.232.136.85:61613/topic/login?id=stomptest&name=test"
	r, e := url.Parse(urlStr)
	if e != nil {
		log.Fatalln("failed to parse the url:" + urlStr)
	}
	h := strings.Split(r.Host, ":")
	fmt.Println("host:", h[0])
	fmt.Println("path:", r.Path)
	fmt.Println("query:", r.Query().Get("ss"))
	fmt.Println("user:", r.User.Username())
	pwd, _ := r.User.Password()
	fmt.Println("pwd:", pwd)
}
