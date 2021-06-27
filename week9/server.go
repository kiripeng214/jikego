package main

import (
	"fmt"
	"kiripeng214/jikego/week9/protocol"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}

	conn, err := listen.Accept()
	defer conn.Close()
	if err != nil {
		return
	}
	packet := protocol.NewDefaultPacket([]byte{})
	tempChanel := make(chan []byte)

	go func() {
		buff := make([]byte, 1024)
		for {
			n, err := conn.Read(buff)
			if err != nil {
				log.Println(err)
				return
			}
			packet.Data = append(packet.Data, buff[:n]...)
			packet.Data = packet.UnPacket(tempChanel)
		}

	}()
	for {
		select {
		case msg := <-tempChanel:
			fmt.Println(string(msg))
		}
	}

}
