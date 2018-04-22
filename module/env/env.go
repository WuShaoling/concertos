package env

import (
	"sync"
	"github.com/concertos/module/util"
	"log"
	"strings"
)

var ENV_ADDR_CONDUCTOR string
var ENV_ADDR_DNS_SERVER string
var ENV_ADDR_NFS_SERVER string
var ENV_ADDR_ETCD_SERVER string
var once sync.Once

const ENV_PATH = "/home/wsl/go/src/github.com/concertos/.config/env.config"

func GetEnv() {

	once.Do(func() {
		if lines, err := util.ReadLines(ENV_PATH); nil != err {
			log.Fatal("read env.config error")
		} else {
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if len(line) != 0 && line[0] != '#' {
					arr := strings.Split(line, "=")
					arr[0] = strings.TrimSpace(arr[0])
					arr[1] = strings.TrimSpace(arr[1])
					switch arr[0] {
					case "ENV_ADDR_CONDUCTOR":
						ENV_ADDR_CONDUCTOR = arr[1]
					case "ENV_ADDR_DNS_SERVER":
						ENV_ADDR_DNS_SERVER = arr[1]
					case "ENV_ADDR_NFS_SERVER":
						ENV_ADDR_NFS_SERVER = arr[1]
					case "ENV_ADDR_ETCD_SERVER":
						ENV_ADDR_ETCD_SERVER = arr[1]
					}
				}
			}
		}
	})
}
