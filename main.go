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
	if *role == "conductor" {
		c := conductor.GetConductor()

		// start manage module
		c.Manager.Start()

		// start rest api module
		c.RestApi.Start()

		// start websocket server module
		c.WebSocket.Start()
	} else if *role == "player" {
		p := player.GetEtcdClient()
		p.Manager.Start()
		p.WebSocket.Start()
	} else {
		log.Fatal("Error args")
	}
}
