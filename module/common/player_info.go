package common

type PlayerInfo struct {
	Id       string   `json:"Id" description:"uniquely identifies'"`
	Ips      []string `json:"Ips" description:"host's ips'"`
	Hostname string   `json:"Hostname" description:"hostname of player"`
	Memory   uint64   `json:"Memory" description:"memory size"`
	Cpu      int      `json:"Cpu" description:"cpu number"`
	State    int      `json:"State" description:"the current status of the player, online/offline"`
}
