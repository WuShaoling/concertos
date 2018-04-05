package main

import (
	"flag"
	"github.com/concertos/conductor"
	"github.com/concertos/player"
	"log"
	"github.com/concertos/module/nfs"
	"github.com/concertos/module/dns"
	"github.com/concertos/network"
)

func main() {

	var role = flag.String("r", "", "conductor | player")
	flag.Parse()

	if *role == "c" {

		log.Println("start cnfs plugin!!")
		go nfs.GetNFSApi().Start()

		log.Println("start cdns plugin!")
		go dns.GetDNSApi().Start()

		c := conductor.GetConductor()

		log.Println("start manager module!")
		go c.Manager.Start()

		log.Println("start websocket module!")
		go c.WebSocket.Start()

		log.Println("start rest api module!")
		c.RestApi.Start()

	} else if *role == "p" {

		log.Println("start cnfs plugin!!")
		go nfs.GetNFSApi().Start()

		p := player.GetPlayer()

		log.Println("start websocket module!")
		go p.WebSocket.Start()

		log.Println("register to conductor!")
		p.Register()

		log.Println("start manager module!")
		p.Manager.Start()

	} else if *role == "t" {
		network.Start()
	} else {
		log.Fatal("Error args")
	}
}
