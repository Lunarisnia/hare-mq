package hare

import (
	"fmt"
	"io"
	"net"
)

type HareServer interface {
	Serve(address string)
	Ping()
}

type HareServerImpl struct {
	connectedClients []net.Conn
}

func NewHareServer() HareServer {
	return &HareServerImpl{
		connectedClients: make([]net.Conn, 0),
	}
}

func (h *HareServerImpl) Serve(address string) {
	listener, err := net.Listen("udp", address)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on: ", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		h.connectedClients = append(h.connectedClients, conn)
		fmt.Println("A new client has connected: ", conn.RemoteAddr().String())
	}
}

func (h *HareServerImpl) Ping() {
	for _, c := range h.connectedClients {
		_, err := io.WriteString(c, "Ping")
		if err != nil {
			panic(err)
		}
		fmt.Println("Pinged: ", c.RemoteAddr().String())
	}
}
