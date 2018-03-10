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
		log.Println("Start manage module")

		// start websocket server module
		go c.WebSocket.Start()
		log.Println("Start websocket module")

		// start rest api module
		c.RestApi.Start()
	} else if *role == "p" {
		p := player.GetPlayer()

		go p.WebSocket.Start()

		log.Println("Player keep alive")
		p.KeepAlive()
	} else {
		log.Fatal("Error args")
	}
}
