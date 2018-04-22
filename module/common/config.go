package common

import (
	"time"
	"github.com/concertos/module/env"
)

const PLAYER_TTL = 18
const PLAYER_HEART_BEAT = 5
const CONTAINER_TTL = 33
const CONTAINER_HEART_BEAT = 10

const (
	PLAYER_STATE_ONLINE  = iota //0
	PLAYER_STATE_OFFLINE        //1
)

const (
	CONTAINER_STATE_NEW      = iota //0
	CONTAINER_STATE_READY           //1
	CONTAINER_STATE_RUNNING         //2
	CONTAINER_STATE_STOPPED         //3
	CONATINER_STATE_WAITTING        //4
	CONATINER_STATE_ERROR           //5
)

const (
	DIAL_TIMEOUT    = 5 * time.Second
	REQUEST_TIMEOUT = 10 * time.Second
)

const RESTAPI_ADDR = ":48080"

const WS_SERVER_POST = ":48081"

func GetWSServerAddr() string {
	return env.ENV_ADDR_CONDUCTOR + WS_SERVER_POST
}

// etcd
const ETCD_PREFIX_CONTAINER_RUNNING = "/container/running/"
const ETCD_PREFIX_CONTAINER_INFO = "/container/info/"

const ETCD_PREFIX_PLAYER_ALIVE = "/player/alive/"
const ETCD_PREFIX_PLAYER_INFO = "/player/info/"

func GetEtcdEndPoints() []string {
	return []string{"http://" + env.ENV_ADDR_ETCD_SERVER + ":2379"}
}

// nfs
const NFS_MOUNT_LOCAL_PATH = "/exports/"
const NFS_MOUNT_CONTAINER_PATH = "/exports/"

func GetNFSMountRemoteAddr() string {
	return env.ENV_ADDR_NFS_SERVER + ":/"
}

// dns
func GetDNSServer_API_ADDR() string {
	return env.ENV_ADDR_DNS_SERVER + ":40001"
}
