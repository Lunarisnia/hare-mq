package main

import (
	"net"
	"time"

	"github.com/Lunarisnia/hare-mq.git/internal/hare"
)

func main() {
	hareServer := hare.NewHareServer()
	go hareServer.Serve(&net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 3000,
	})
	for {
		time.Sleep(2 * time.Second)
		hareServer.Ping()
	}
}
