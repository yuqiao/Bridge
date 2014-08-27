package stomp

import (
	"errors"
	"github.com/gmallard/stompngo"
	"github.com/golang/glog"
	"net"
	"net/url"
	"strings"
)

type StompInfo struct {
	id     string
	host   string
	port   string
	user   string
	passwd string
	topic  string
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

	h = strings.Split(u.Host, ":")
	info.host = h[0]
	if len(h) == 2 {
		info.port = h[1]
	} else {
		info.port = "61613"
	}

	info.user = u.User.Username()
	info.passwd = u.User.Password()
	info.id = u.Query().Get("id")
	info.topic = u.Path
}

func (s *StompInfo) String() string {
	return "[StompInfo]"
}

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
