package network

import (
	"github.com/concertos/player/util"
	"fmt"
	"strconv"
)

func SetIP() bool {
	files := ReadDir("/var/run/netns")
	if files == nil {
		fmt.Println("files is nil, mkdir /var/run/netns")
		return false
	}

	/*
	# 创建 veth pair，并把一端加到网桥上
	ip link add veth0 type veth peer name veth1
	ip link set dev veth0 master br0
	ip link set dev veth0 up

	# 配置容器内部的网络和 IP
	ip link set dev veth1 netns container1
	ip netns exec container1 ip link set veth1 name eth0
	ip netns exec container1 ip addr add 10.20.1.2/24 dev eth0
	ip netns exec container1 ip link set eth0 up
	*/

	count := 0

	for k, v := range files {
		IP := "10.0.1.2" + strconv.Itoa(k) + "/24"
		veth0 := "veth" + strconv.Itoa(count)
		count++
		veth1 := "veth" + strconv.Itoa(count)
		count++
		util.ExecShell("ip link add " + veth0 + " type veth peer name " + veth1)
		util.ExecShell("ip link set dev " + veth0 + " master " + BRIDGE_NAME)
		util.ExecShell("ip link set dev " + veth0 + " up")

		util.ExecShell("ip link set dev " + veth1 + " netns " + v)
		util.ExecShell("ip netns exec " + v + " ip link set " + veth1 + " name eth0")
		util.ExecShell("ip netns exec " + v + " ip addr add " + IP + " dev eth0")
		util.ExecShell("ip netns exec " + v + " ip link set eth0 up")
	}

	return true
}
