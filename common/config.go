package common

import "time"

const NAMESPACE = "cos"
const TTL = 15
const HEART_BEAT = 5

const (
	PLAYER_STATE = iota //0
	ONLINE              //1
	OFFLINE             //2
	UNKONWEN            //3
)

const (
	CONTAINER_TYPE = iota //0
	REG                   //1
	DOCKER                //2
)

const (
	CONTAINER_STATE = iota //0
	RUNNING                //1
	STOPPED                //2
	PAUSED                 //3
)

const ETCD_PREFIX_PLAYERS_ALIVE = "/players/alive/"
const ETCD_PREFIX_PLAYERS_INFO = "/players/info/"
const ETCD_PREFIX_USERS_ALIVE = "/users/alive/"
const ETCD_PREFIX_USERS_INFO = "/users/info/"

func GetEtcdPoints() []string {
	return []string{"http://127.0.0.1:2379"}
}

const (
	DIAL_TIMEOUT    = 5 * time.Second
	REQUEST_TIMEOUT = 10 * time.Second
)

const (
	ENTITY_TYPE              = iota //0
	ENTITY_PLAYER                   //1
	ENTITY_CONTAINER               //2
	ENTITY_USER                     //3
)
