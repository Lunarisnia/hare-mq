package hare

import (
	"fmt"
	"io"
	"net"
)

type HareServer interface {
	Serve(tcpAddr *net.TCPAddr)
	Ping()
	ReadMessage()
}

type HareServerImpl struct {
	connectedClients []net.Conn
}

func NewHareServer() HareServer {
	return &HareServerImpl{
		connectedClients: make([]net.Conn, 0),
	}
}

func (h *HareServerImpl) Serve(tcpAddr *net.TCPAddr) {
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Printf("Listening on: %v:%v\n", tcpAddr.IP, tcpAddr.Port)

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
			c.Close()
		}
		fmt.Println("Pinged: ", c.RemoteAddr().String())
	}
}

func (h *HareServerImpl) ReadMessage() {
	body := make([]byte, 32)
	for _, c := range h.connectedClients {
		_, err := c.Read(body)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%v say: %v\n", c.RemoteAddr().String(), string(body))
		fmt.Fprint(c, "Pong")
	}
}
