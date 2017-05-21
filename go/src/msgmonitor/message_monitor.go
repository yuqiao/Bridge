package main

import (
	"alilib/stomp"
	"encoding/json"
	"fmt"
	"github.com/gmallard/stompngo"
	"log"
)

func main() {
	info, _ := stomp.GetStompInfo("mq://wangxin:wangxin@10.232.136.85:61613/topic/logon")
	out, _ := json.Marshal(info)
	fmt.Printf("info: %s\n", out)

	sconn, _ := stomp.GetStompConn("mq://wangxin:wangxin@10.232.136.85:61613/topic/logon", "gotest")

	r := sconn.Subscribe()
	defer sconn.Unsubscribe()

	for i := 1; i <= 2; i++ {
		m := <-r
		if m.Error != nil {
			log.Fatalln(m.Error)
		}
		fmt.Print("Frame Type: %s\n", m.Message.Command)
		if m.Message.Command != stompngo.MESSAGE {
			fmt.Println("do not handle the message")
			continue
		}
		h := m.Message.Headers
		for j := 0; j < len(h)-1; j += 2 {
			fmt.Printf("Header: %s:%s", h[j], h[j+1])
		}
		fmt.Printf("Payload: %s\n", string(m.Message.Body))
	}
}
