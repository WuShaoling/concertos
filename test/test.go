package main;

import (
	"flag"
	"log"
)

func main() {
	var addr1 = flag.String("addr", "localhost:8081", "http service address")
	log.Println(*addr1)
}
