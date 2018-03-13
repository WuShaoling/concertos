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

		//// mount to nfs server
		//nfs := nfs.GetNFSApi()
		//nfs.UMount()
		//nfs.Mount()

		c := conductor.GetConductor()

		c.Manager.Start()

		go c.WebSocket.Start()

		c.RestApi.Start()

	} else if *role == "p" {

		p := player.GetPlayer()

		go p.WebSocket.Start()

		p.Register()

		p.KeepAlive()

	} else {
		log.Fatal("Error args")
	}
}
