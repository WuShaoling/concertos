package restapi

import "github.com/concertos/common"

type ContainerResource struct {
	myEtcdClient *common.MyEtcdClient
}

func GetContainerResource() *ContainerResource {
	return &ContainerResource{
		myEtcdClient: common.GetMyEtcdClient(),
	}
}
