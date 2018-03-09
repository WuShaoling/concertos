package common

import "time"

const NAMESPACE = "cos"
const TTL = 20
const HEART_BEAT = 5

const (
	PLAYER_STATE         = iota //0
	PLAYER_STATE_ONLINE         //1
	PLAYER_STATE_OFFLINE        //2
)

const (
	CONTAINER_TYPE = iota //0
	REG                   //1
	DOCKER                //2
)

const (
	CONTAINER_STATE         = iota //0
	CONTAINER_STATE_RUNNING        //1
	CONTAINER_STATE_STOPPED        //2
)

// common prefix
const ETCD_PREFIX_CONATINER_WAIT_START = "/container/wait-start/"
const ETCD_PREFIX_CONTAINER_ALIVE = "/container/alive/"
const ETCD_PREFIX_CONTAINER_INFO = "/container/info/"
const ETCD_PREFIX_CONTAINER_SHOULD_ALIVE = "/container/should-alive/"

const ETCD_PREFIX_PLAYER_ALIVE = "/player/alive/"
const ETCD_PREFIX_PLAYER_INFO = "/player/info/"

const ETCD_PREFIX_USER_ALIVE = "/user/alive/"
const ETCD_PREFIX_USER_INFO = "/user/info/"

func GetEtcdPoints() []string {
	return []string{"http://127.0.0.1:2379"}
}

const (
	DIAL_TIMEOUT    = 5 * time.Second
	REQUEST_TIMEOUT = 10 * time.Second
)

const (
	ENTITY_TYPE      = iota //0
	ENTITY_PLAYER           //1
	ENTITY_CONTAINER        //2
	ENTITY_USER             //3
)
