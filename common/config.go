package common

const NAMESPACE = "cos"
const TTL = 30
const HEART_BEAT = 10

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

const ETCD_PREFIX_PLAYERS_INFO = "/players/"
const ETCD_PREFIX_USERS_INFO = "/users/"

func GetEtcdPoints() []string {
	return []string{"http://127.0.0.1:2379"}
}
