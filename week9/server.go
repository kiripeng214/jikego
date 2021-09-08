package main

import (
	"fmt"
	"io"
	"kiripeng214/jikego/week9/xconn"
)

func main() {
	//listen, err := net.Listen("tcp", ":8080")
	//if err != nil {
	//	return
	//}
	xcon := xconn.Conn{
		NetWork: "tcp",
		Address: "127.0.0.1",
		Port:    "8080",
	}
	xcon.ServerInit()
	defer xcon.ServerClose()
	for {
		read, err := xcon.Read()
		if err == io.EOF {
			xcon.ServerInit()
		} else if err != nil {
			fmt.Println(err)
			return
		}

		if read != nil {
			fmt.Println(string(read.Data))
		}

	}
}
