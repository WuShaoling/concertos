package main

import (
	"flag"
	"fmt"
	"github.com/concertos/conductor"
	"github.com/concertos/player"
)

func main() {
	var role = flag.String("role", "", "master | worker")
	flag.Parse()
	if *role == "etcd" {
		conductor.StartConductor()
	} else if *role == "player" {
		player.StartPlayer()
	} else {
		fmt.Println("Error args")
	}
}
