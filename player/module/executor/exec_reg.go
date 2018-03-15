package executor

import (
	"github.com/concertos/module/common"
	"github.com/concertos/module/entity"
	"github.com/concertos/player/util"
	"github.com/concertos/player/module/manager"
)

func (d *RegExecutor) Install() error {
	return nil;
}

func (d *RegExecutor) Start(con *entity.ContainerInfo) string {
	// put to running info
	d.myEtcdClient.Put(common.ETCD_PREFIX_CONTAINER_RUNNING+con.Id, con.Id)

	// update container state and put to etcd
	con.State = common.CONTAINER_STATE_RUNNING
	d.myEtcdClient.Put(common.ETCD_PREFIX_CONTAINER_INFO+con.Id, string(util.MyJsonMarshal(con)))

	//register to manager module, keep alive
	manager.GetContainerManager().Register(con.Id)

	// config network

	// start reg

	return "ok";
}

func (d *RegExecutor) Stop() error {
	return nil;
}

func (d *RegExecutor) Remove() error {
	return nil;
}

type RegExecutor struct {
	myEtcdClient *common.MyEtcdClient
}

func GetRegExecutor() *RegExecutor {
	return &RegExecutor{
		myEtcdClient: common.GetMyEtcdClient(),
	}
}
