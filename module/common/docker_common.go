package common

import (
	"sync"
	"github.com/docker/docker/client"
	"log"
)

var once3 sync.Once
var Cli *client.Client

func GetDockerClient() *client.Client {
	once3.Do(func() {
		cli, err := client.NewEnvClient()
		if err != nil {
			log.Println(err)
		}
		Cli = cli
	})
	return Cli
}
