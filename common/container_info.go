package common

type ContainerInfo struct {
	Id          string   `json:"Id" description:"uniquely identifies of container"`
	Name        string   `json:"Name" description:"name of container'"`
	User        string   `json:"UserInfo" description:"username of container"`
	Ips         []string `json:"Ips" description:"container's ips"`
	Command     string   `json:"Ips" description:"command"`
	State       int      `json:"State" description:"the current status of the container, running, stopped, paused..."`
	Discribe    string   `json:"Discribe" description:"additional description information"`
	Created     string   `json:"Created" description:"created time"`
	LastStopped string   `json:"LastStopped" description:"last stopped time"`
}
