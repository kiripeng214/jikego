package main

import (
	"fmt"
	"kiripeng214/jikego/week9/protocol"
	"log"
	"net"
	"time"
)

func sender(conn net.Conn) {

}

func main() {

	server := "127.0.0.1:8080"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)

	if err != nil {
		log.Println(err)
		return

	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		log.Println(err)
		return

	}

	defer conn.Close()

	fmt.Println("connect success")

	packet := protocol.NewDefaultPacket([]byte{})
	for i := 0; i < 1000; i++ {
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
		packet.Data = []byte(words)
		fmt.Println(string(packet.Data))
		conn.Write(packet.Packet())

	}

	fmt.Println("send over")
	time.Sleep(20 * time.Second)

}
