package scheduler

import (
	"github.com/concertos/module/common"
	"github.com/coreos/etcd/clientv3"
	"time"
	"math/rand"
	"log"
)

type RandomAlgorithm struct {
	myEtcdClient *common.MyEtcdClient
}

func GetRandomAlgorithm() *RandomAlgorithm {
	return &RandomAlgorithm{
		myEtcdClient: common.GetMyEtcdClient(),
	}
}

func (ra *RandomAlgorithm) GetPlayerId() string {
	resp := ra.myEtcdClient.Get(common.ETCD_PREFIX_PLAYER_INFO, clientv3.WithPrefix())
	players := *ra.myEtcdClient.ConvertToPlayerInfo(resp)
	if len(players) <= 0 {
		log.Println("No player alive")
		return ""
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	log.Println("Select player : ", players[r.Intn(len(players))].Id)
	return players[r.Intn(len(players))].Id
}
