package stomp

import (
	"errors"
	"fmt"
	"github.com/gmallard/stompngo"
	"github.com/golang/glog"
	"log"
	"net"
	"net/url"
	"strings"
)

type StompInfo struct {
	Id     string
	Host   string
	Port   string
	User   string
	Passwd string
	Topic  string
}

type StompConn struct {
	StompInfo

	conn *stompngo.Connection
}

func GetStompInfo(urlStr string) (info StompInfo, e error) {
	info = StompInfo{}
	u, e := url.Parse(urlStr)
	if e != nil {
		return
	}

	if u.Scheme != "mq" {
		glog.Error("invalid scheme:" + u.Scheme)
		e = errors.New("invalid scheme:" + u.Scheme)
		return
	}

	h := strings.Split(u.Host, ":")
	info.Host = h[0]
	if len(h) == 2 {
		info.Port = h[1]
	} else {
		info.Port = "61613"
	}

	info.User = u.User.Username()
	info.Passwd, _ = u.User.Password()
	info.Id = u.Query().Get("id")
	info.Topic = u.Path

	return info, nil
}

func GetStompConn(urlStr, id string) (sconn StompConn, e error) {
	sconn = StompConn{}
	sconn.StompInfo, e = GetStompInfo(urlStr)
	if e != nil {
		log.Fatalln(e)
	}

	n, e := net.Dial("tcp", net.JoinHostPort(sconn.Host, sconn.Port))
	if e != nil {
		log.Fatalln(e)
	}

	fmt.Println(sconn.Id + " dail complete .")
	ch := ConnectHeaders(sconn.User, sconn.Passwd, sconn.Host)
	sconn.conn, e = stompngo.Connect(n, ch)
	sconn.Id = id
	return sconn, e
}

func (s *StompInfo) String() string {
	return fmt.Sprintf("[StompInfo]{host=%s, port=%s id=%s")
}

func (sconn *StompConn) Subscribe() <-chan stompngo.MessageData {
	return subscribe(sconn.conn, sconn.Topic, sconn.Id)
}

func (sconn *StompConn) Unsubscribe() {
	unsubscribe(sconn.conn, sconn.Id)
}

func ConnectHeaders(user, passwd, hst string) stompngo.Headers {
	h := stompngo.Headers{}
	h = h.Add("login", user)
	h = h.Add("passcode", passwd)
	h = h.Add("accept-version", stompngo.SPL_11).Add("host", hst)
	return h
}

func subscribe(c *stompngo.Connection, dest, id string) <-chan stompngo.MessageData {
	h := stompngo.Headers{"destination", dest, "ack", "auto"}

	h = h.Add("id", id)
	r, e := c.Subscribe(h)
	if e != nil {
		log.Fatalln("subscribe failed", e)
	}
	return r
}

func unsubscribe(c *stompngo.Connection, id string) {
	h := stompngo.Headers{}

	h = h.Add("id", id)
	e := c.Unsubscribe(h)
	if e != nil {
		log.Fatalln("unsubscribe failed.", e)
	}
	return
}
