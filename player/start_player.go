package player

import (
	"github.com/concertos/player/module/etcd"
)

func StartPlayer() {
	etcdClient := etcd.GetEtcdClient()

	etcdClient.HeartBeat()
}
