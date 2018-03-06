package common

type UserInfo struct {
	Id       string `json:"id" description:"identifier of the user"`
	Password string `json:"password" description:"password of user"`
	Name     string `json:"name" description:"name of the user"`
	Created  int    `json:"created" description:"created time`
}
