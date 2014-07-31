package ipc

import (
	"fmt"
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(request, params string) *Response {

	code := "200"
	body := "ECHO:" + request + ", " + params
	return &Response{code, body}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {

	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, e1 := client1.Call("From Client1", "1")
	resp2, e2 := client2.Call("From Client2", "2")

	if e1 != nil || e2 != nil {
		t.Error("IpcCall.Call failed. e1:", e1, "e2:", e2)
	}

	fmt.Println(resp1)
	fmt.Println(resp2)
	if resp1.Body != "ECHO" {
		t.Error("IpcCall.Call failed. resp1:", resp1, "rsp2:", resp2)
	}

	client1.Close()
	client2.Close()
}
