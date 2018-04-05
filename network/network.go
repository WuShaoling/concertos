package network

import (
	"github.com/concertos/player/util"
	"log"
	"strings"
	"os"
)

const FLANNEL_ENV_APTH = "/run/flannel/subnet.env"
const DOCKER_CONFIG = "/etc/docker/daemon.json"

func Start() {
	startFlannel()
	configDocker()
}

func startFlannel() {
	cmd := "sudo nohup flanneld > /home/wsl/conceros/flannel.log 2>&1 &"
	if res, err := util.ExecShell(cmd); err != nil {
		log.Fatal("start flannel error", err)
	} else if len(res) != 0 {
		log.Fatal("start flannel res: ", res)
	} else {
		log.Println("start flannel success")
	}
}

func configDocker() {

	if lines, err := util.ReadLines(FLANNEL_ENV_APTH); nil != err {
		log.Fatal(err)
	} else {
		docker := "{\"registry-mirrors\": [\"https://258g3pnb.mirror.aliyuncs.com\"],\"bip\": \""
		docker_s2 := "\",\"mtu\": 1450,\"ip-masq\":true}"

		for _, line := range lines {
			if strings.Contains(line, "FLANNEL_SUBNET") {
				arr := strings.Split(line, "=")
				docker = docker + strings.Replace(arr[1], "\n", "", -1) + docker_s2
				break
			}
		}

		// delete file
		if err = os.Remove(DOCKER_CONFIG); nil != err {
			log.Println(err)
		}
		// write new config
		util.WriteFile(docker, DOCKER_CONFIG)
		cmd := "sudo service docker restart"
		util.ExecShell(cmd)
	}
}
