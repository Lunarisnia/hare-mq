package hare

import (
	"fmt"
	"net"
)

type HareClient interface {
	Connect(serverAddress string)
	Ping()
	Close()
	ReadMessage()
}

type HareClientImpl struct {
	serverConnection net.Conn
}

func NewHareClient() HareClient {
	return &HareClientImpl{}
}

func (h *HareClientImpl) Connect(serverAddress string) {
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		panic(err)
	}
	h.serverConnection = conn
}

func (h *HareClientImpl) Ping() {
	_, err := h.serverConnection.Write([]byte("Ping"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Pinged server")
}

func (h *HareClientImpl) Close() {
	err := h.serverConnection.Close()
	if err != nil {
		panic(err)
	}
}

func (h *HareClientImpl) ReadMessage() {
	body := make([]byte, 32)
	for {
		_, err := h.serverConnection.Read(body)
		if err != nil {
			panic(err)
		}
		fmt.Println("Server said: ", string(body))
	}
}
