package config

const NameSpace = "cos"
const TTL = 60
const HEARTBEAT = 20
const (
	CONTAINER_TYPE = iota
	REG
	DOCKER
)

const (
	RUNNING_STATE = iota
	RUNNING
	STOPPED
	PAUSE
)

func GetEtcdPoints() []string {
	return []string{"http://127.0.0.1:2379"}
}
