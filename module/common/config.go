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
	CONATINER_STATE_ERROR           //5
)

const (
	DIAL_TIMEOUT    = 5 * time.Second
	REQUEST_TIMEOUT = 10 * time.Second
)

const RESTAPI_ADDR = ":8080"

const WS_SERVER_ADDR = "localhost:8081"

// etcd
const ETCD_PREFIX_CONTAINER_RUNNING = "/container/running/"
const ETCD_PREFIX_CONTAINER_INFO = "/container/info/"

const ETCD_PREFIX_PLAYER_ALIVE = "/player/alive/"
const ETCD_PREFIX_PLAYER_INFO = "/player/info/"

const ETCD_PREFIX_USER_ALIVE = "/user/alive/"
const ETCD_PREFIX_USER_INFO = "/user/info/"

func GetEtcdPoints() []string {
	return []string{"http://127.0.0.1:2379"}
}

// nfs
const NFS_SERVER_ADDR = "115.159.30.115"
const NFS_MOUNT_REMOTE_ADDR = "115.159.30.115:/nfs/data/"
const NFS_MOUNT_LOCAL_PATH = "/exports/"

// dns
const DNS_SERVER_ADDR = "localhost:53"
const DNS_SERVER_API_ADDR = "localhost:8082"
