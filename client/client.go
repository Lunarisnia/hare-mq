package main

import (
	"github.com/Lunarisnia/hare-mq.git/internal/hare"
)

func main() {
	hareClient := hare.NewHareClient()
	hareClient.Connect("127.0.0.1:3000")
	hareClient.ReadMessage()
}
