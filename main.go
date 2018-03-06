package main

import (
	"flag"
	"fmt"
	"github.com/concertos/conductor"
	"github.com/concertos/player"
)

func main() {
	var role = flag.String("role", "", "conductor | player")
	flag.Parse()
	if *role == "conductor" {
		conductor.StartConductor()
	} else if *role == "player" {
		player.StartPlayer()
	} else {
		fmt.Println("Error args")
	}
}
