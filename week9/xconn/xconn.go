package xconn

import (
	"fmt"
	"io"
	"kiripeng214/jikego/week9/message"
	"kiripeng214/jikego/week9/protocol"
	"net"
)

type XMsg interface {
	Read() (*message.Message, error)
	Write(*message.Message) error
}

var _ XMsg = (*Conn)(nil)

type Conn struct {
	NetWork   string
	Address   string
	Port      string
	server    net.Conn
	client    *net.TCPConn
	readPack  *protocol.Packet
	writePack *protocol.Packet
}

//func (c *Conn) init()  {
//	c.packet = protocol.NewDefaultPacket([]byte{})
//}

func (c *Conn) ClientInt() (err error) {
	server := fmt.Sprintf("%v:%v", c.Address, c.Port)
	c.writePack = protocol.NewDefaultPacket([]byte{})

	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		return err
	}
	c.client, err = net.DialTCP("tcp", nil, tcpAddr)
	return err
}

func (c *Conn) ServerInit() (err error) {
	c.readPack = protocol.NewDefaultPacket([]byte{})
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	c.server, err = listen.Accept()
	return err
}

func (c *Conn) ServerClose() (err error) {

	return c.server.Close()
}

func (c *Conn) ClientClose() (err error) {
	return c.client.Close()
}

func (c *Conn) Read() (*message.Message, error) {
	var whole []byte
	var buff []byte
	m := &message.Message{}
	if len(c.readPack.Data) != 0 {
		c.readPack.Data, whole = c.readPack.UnPacket()
		if len(whole) != 0 {
			m.Data = whole
			return m, nil
		}
	}
	for {
		buff = make([]byte, 1024)
		n, err := c.server.Read(buff)
		if err != nil {
			return nil, err
		}
		c.readPack.Data = append(c.readPack.Data, buff[:n]...)
		c.readPack.Data, whole = c.readPack.UnPacket()
		if len(whole) != 0 {
			m.Data = whole
			return m, nil
		}
		if len(whole) == 0 && len(c.readPack.Data) == 0 {
			return nil, io.EOF
		}
	}
}

func (c *Conn) Write(m *message.Message) error {
	c.writePack.Data = m.Data
	_, err := c.client.Write(c.writePack.Packet())
	return err
}
