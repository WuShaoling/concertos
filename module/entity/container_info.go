package entity

type ContainerInfo struct {
	Id          string `json:"Id" description:"uniquely identifies of container"`
	Name        string `json:"Name" description:"name of container'"`
	User        string `json:"UserInfo" description:"userid of container"`
	Ip          string `json:"Ips" description:"container's ip"`
	Command     string `json:"Ips" description:"command"`
	State       int    `json:"State" description:"the current status of the container, running, stopped, paused..."`
	Discribe    string `json:"Discribe" description:"additional description information"`
	Created     int64  `json:"Created" description:"created time"`
	LastStopped int64  `json:"LastStopped" description:"last stopped time"`
	PlayerId    string `json:"PlayerId" description:""`
}

type Docker struct {
}

type Reg struct {
}
