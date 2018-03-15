package common

import "time"

const NAMESPACE = "cos"
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
)

const (
	DIAL_TIMEOUT    = 5 * time.Second
	REQUEST_TIMEOUT = 10 * time.Second
)

const ETCD_PREFIX_CONTAINER_RUNNING = "/container/running/"
const ETCD_PREFIX_CONTAINER_INFO = "/container/new/"

const ETCD_PREFIX_PLAYER_ALIVE = "/player/alive/"
const ETCD_PREFIX_PLAYER_INFO = "/player/info/"

const ETCD_PREFIX_USER_ALIVE = "/user/alive/"
const ETCD_PREFIX_USER_INFO = "/user/info/"

const NFS_MOUNT_PATH = "/nfs/share/"

const RESTAPI_ADDR = ":8080"

func GetEtcdPoints() []string {
	return []string{"http://127.0.0.1:2379"}
}

func GetWebSocketPort() string {
	return "8081"
}

func GetWebSocketServerAddress() string {
	return "localhost:8081"
}

func GetNFSServerAddress() string {
	return "192.168.1.149"
}
