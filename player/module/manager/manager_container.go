package manager

import (
	"github.com/concertos/module/common"
	"github.com/coreos/etcd/clientv3"
	"context"
	"log"
)

func (cm *ContainerManager) WatchRunningContainer() {
	rch := cm.myEctdClient.GetClientV3().Watch(context.Background(), common.ETCD_PREFIX_CONTAINER_RUNNING, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			log.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			switch  ev.Type.String() {
			case "DELETE":
			default:
			}
		}
	}
}

type ContainerManager struct {
	myEctdClient *common.MyEtcdClient
}

func GetManagerContainer() *ContainerManager {
	return &ContainerManager{
		myEctdClient: common.GetMyEtcdClient(),
	}
}
