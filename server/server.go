package main

import "github.com/Lunarisnia/hare-mq.git/internal/hare"

func main() {
	hareServer := hare.NewHareServer()
	hareServer.Serve("0.0.0.0:3000")
}
