package entity

type ContainerInfo struct {
	Id        string
	Name      string
	Ip        string
	Command   string
	Describe  string
	PlayerId  string
	BaseImage string
	State     int
	CPU       int
	Memory    uint64
	Port      uint32
}
