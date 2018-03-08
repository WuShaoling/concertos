package restapi

import "github.com/concertos/common"

type PlayerResource struct {
	myEtcdClient *common.MyEtcdClient
}

func GetPlayerResource() *PlayerResource {
	return &PlayerResource{
		myEtcdClient: common.GetMyEtcdClient(),
	}
}
