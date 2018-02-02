package osnetwork

import (
	"github.com/concertos/player/util"
	"fmt"
)

/**
setup bridge/vxlan, and connect vxlan to bridge
 */
func ConfigNetwork(host_eth string) bool {
	_, err := util.ExecShell("ip link delete " + BRIDGE_NAME)
	_, err = util.ExecShell("ip link delete " + VXLAN_NAME)

	_, err = util.ExecShell("ip link add " + VXLAN_NAME + " type vxlan id " +
		VXLAN_ID + " dstport " + VXLAN_PORT + " group " + VXLAN_GROUP + " dev " + host_eth)
	if err != nil {
		fmt.Println("error!!! ", err, "add vxlan failed")
		return false
	}

	_, err = util.ExecShell("ip addr add " + VXLAN_IP + " dev " + VXLAN_NAME)
	if err != nil {
		fmt.Println("error!!! ", err, "set vxlan ip failed")
		return false
	}

	_, err = util.ExecShell("ip link add " + BRIDGE_NAME + " type bridge")
	if err != nil {
		fmt.Println("error!!! ", err, "add bridge failed")
		return false
	}

	_, err = util.ExecShell("ip link set " + VXLAN_NAME + " master " + BRIDGE_NAME)
	if err != nil {
		fmt.Println("error!!! ", err, "bind failed")
		return false
	}

	_, err = util.ExecShell("ip link set " + VXLAN_NAME + " up")
	if err != nil {
		fmt.Println("error!!! ", err, "set vxlan up failed")
		return false
	}

	_, err = util.ExecShell("ip link set " + BRIDGE_NAME + " up")
	if err != nil {
		fmt.Println("error!!! ", err, "set bridge up failed")
		return false
	}
	return true
}

const BRIDGE_NAME = "cos_br0"
const VXLAN_NAME = "cos_vxlan0"
const VXLAN_ID = "101"
const VXLAN_GROUP = "239.1.1.1"
const VXLAN_PORT = "4789"
const VXLAN_IP = "10.0.1.1/24"
