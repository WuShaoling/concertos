package docker

import (
	"github.com/concertos/module/util"
	"log"
	"strings"
	"os"
	"github.com/concertos/module/env"
)

const FLANNEL_ENV_APTH = "/run/flannel/subnet.env"
const DOCKER_CONFIG = "/etc/docker/daemon.json"

func Config() {
	log.Println("config docker ...")

	lines, err := util.ReadLines(FLANNEL_ENV_APTH)
	if nil != err {
		log.Fatal(err)
	}
	docker := "{\"registry-mirrors\": [\"https://258g3pnb.mirror.aliyuncs.com\"],\"bip\": \""
	docker_s2 := "\",\"mtu\": 1450,\"ip-masq\":true, \"dns\": [\"" + env.ENV_ADDR_DNS_SERVER + "\"]}"

	for _, line := range lines {
		if strings.Contains(line, "FLANNEL_SUBNET") {
			arr := strings.Split(line, "=")
			docker = docker + strings.Replace(arr[1], "\n", "", -1) + docker_s2
			break
		}
	}

	if err = os.Remove(DOCKER_CONFIG); nil != err {
		log.Println(err)
	}

	util.WriteFile(docker, DOCKER_CONFIG)
	util.ExecShell("sudo service docker restart")
}
