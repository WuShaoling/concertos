package manager

import (
	"github.com/concertos/module/common"
	"github.com/coreos/etcd/clientv3"
	"context"
	"log"
	"os"
	"os/signal"
	"time"
)

func (cm *ContainerManager) KeepContainerAlive() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ticker := time.NewTicker(time.Second * common.CONTAINER_HEART_BEAT)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for k, v := range cm.Containers {
				key := common.ETCD_PREFIX_CONTAINER_RUNNING + k
				resp, err := cm.myEtcdClient.GetClientV3().Grant(context.TODO(), common.CONTAINER_TTL)
				if err != nil {
					log.Println(err)
				} else if _, err = cm.myEtcdClient.GetClientV3().Put(context.TODO(), key, v, clientv3.WithLease(resp.ID)); nil != err {
					log.Println(err)
				}
			}
		case <-interrupt:
			log.Println("System interrupt, keep container alive interrupt, ticker stop")
			return
		}
	}
}

func (cm *ContainerManager) Remove(id string) {
	delete(cm.Containers, id)
}

func (cm *ContainerManager) Register(id string) {
	log.Println("register container ", id)
	cm.Containers[id] = id
}

type ContainerManager struct {
	myEtcdClient *common.MyEtcdClient
	Containers   map[string](string)
}

func GetContainerManager() *ContainerManager {
	return &ContainerManager{
		myEtcdClient: common.GetMyEtcdClient(),
		Containers:   make(map[string](string)),
	}
}
