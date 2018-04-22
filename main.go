package main

import (
	"flag"
	"log"
	"github.com/concertos/module/env"
	"github.com/concertos/conductor"
	"github.com/concertos/player"
	"github.com/concertos/module/dns"
	"fmt"
)

func main() {

	//读取环境变量
	log.Println("get env")
	env.GetEnv()

	var role = flag.String("r", "", "conductor | player | test")
	flag.Parse()

	if *role == "c" || *role == "p" {

		//// 修改 docker 参数并重启 docker
		//docker.Config()

		////挂载nfs
		//log.Println("start nfsd plugin!!")
		//nfs.GetNFSApi().Start()

		if *role == "c" {

			c := conductor.GetConductor()

			log.Println("start websocket module!")
			go c.WebSocket.Start()

			log.Println("start rest api module!")
			c.RestApi.Start()

		} else if *role == "p" {
			p := player.GetPlayer()

			log.Println("start websocket module!")
			go p.WebSocket.Start()

			log.Println("register to conductor!")
			p.Register()

			log.Println("start manager module!")
			p.Manager.Start()

		}
	}

	if *role == "t" {
		fmt.Println(dns.GetDNSApi().GetAll())
	} else {
		log.Fatal("Error args")
	}
}
