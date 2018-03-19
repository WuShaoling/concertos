package entity

type ContainerInfo struct {
	Id        string `json:"Id" description:"uniquely identifies of container"`
	Name      string `json:"Name" description:"name of container'"`
	User      string `json:"User" description:"user id of container"`
	Ip        string `json:"Ip" description:"container's ip"`
	Command   string `json:"Command" description:"command"`
	Describe  string `json:"Describe" description:"additional description information"`
	PlayerId  string `json:"PlayerId" description:""`
	BaseImage string `json:"BaseImage" description:""`
	State     int    `json:"State" description:"the current status of the container, running, stopped, paused..."`
	Created   int64  `json:"Created" description:"created time"`
}
