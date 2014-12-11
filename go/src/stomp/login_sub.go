package main

import (
	"fmt"
	"github.com/gmallard/stompngo"
	"log"
	"net"
)

const (
	stompid = "gotest"
	host    = "10.232.136.85"
	port    = "61613"
	user    = "wangxin"
	passwd  = "wangxin"
)

func ConnectHeaders(user, passwd, hst string) stompngo.Headers {
	h := stompngo.Headers{}
	h = h.Add("login", user)
	h = h.Add("passcode", passwd)
	h = h.Add("accept-version", stompngo.SPL_11).Add("host", hst)
	return h
}

func Subscribe(c *stompngo.Connection, d, i, a string) <-chan stompngo.MessageData {
	h := stompngo.Headers{"destination", d, "ack", "auto"}

	h = h.Add("id", stompid)
	r, e := c.Subscribe(h)
	if e != nil {
		log.Fatalln("subscribe failed", e)
	}
	return r
}

func Unsubscribe(c *stompngo.Connection, d, i string) {
	h := stompngo.Headers{}

	h = h.Add("id", stompid)
	e := c.Unsubscribe(h)
	if e != nil {
		log.Fatalln("unsubscribe failed.", e)
	}
	return
}

func main() {
	n, e := net.Dial("tcp", net.JoinHostPort(host, port))
	if e != nil {
		log.Fatalln(e)
	}

	fmt.Println(stompid + " dail complete .")
	ch := ConnectHeaders()
	conn, e := stompngo.Connect(n, ch)

	if e != nil {
		log.Fatalln("stomp connect failed!", e)
	}

	fmt.Println(stompid + " stomp connect complete ...")
	fmt.Println(stompid+" connectd headers ", conn.ConnectResponse.Headers)

	d := "/topic/logon"
	i := stompngo.Uuid()

	r := Subscribe(conn, d, i, "auto")

	for i := 1; i <= 10; i++ {
		m := <-r
		fmt.Println(stompid + " channel read complete ...")
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

	Unsubscribe(conn, d, i)
	fmt.Println(stompid + " unsubscribe complete")

	e = conn.Disconnect(stompngo.Headers{})

	fmt.Println(stompid + " stomp disconnect complete")

	e = n.Close()
	fmt.Println(stompid + " network close complete")
}
