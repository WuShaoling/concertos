package main

import (
	"flag"
	"github.com/concertos/conductor"
	"github.com/concertos/player"
	"log"
)

func main() {
	var role = flag.String("r", "", "conductor | player")
	flag.Parse()
	if *role == "c" {
		c := conductor.GetConductor()

		// start manage module
		c.Manager.Start()

		// start websocket server module
		go c.WebSocket.Start()

		// start rest api module
		c.RestApi.Start()
	} else if *role == "p" {
		p := player.GetPlayer()

		// start web socket goroutine
		go p.WebSocket.Start()

		// register player
		p.Register()

		// start keep player alive goroutine
		p.KeepAlive()
	} else {
		log.Fatal("Error args")
	}
}
