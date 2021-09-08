package main

import (
	"fmt"
	"kiripeng214/jikego/week9/message"
	"kiripeng214/jikego/week9/protocol"
	"kiripeng214/jikego/week9/xconn"
	"time"
)

func main() {

	xcon := xconn.Conn{
		NetWork: "tcp",
		Address: "127.0.0.1",
		Port:    "8080",
	}
	xcon.ClientInt()
	defer xcon.ClientClose()
	packet := protocol.NewDefaultPacket([]byte{})
	for i := 0; i < 1000; i++ {
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
		packet.Data = []byte(words)
		xcon.Write(&message.Message{Data: packet.Data})
	}

	fmt.Println("send over")
	time.Sleep(20 * time.Second)

}
