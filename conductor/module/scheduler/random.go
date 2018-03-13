package scheduler

import (
	"github.com/concertos/module/common"
	"github.com/coreos/etcd/clientv3"
	"errors"
	"time"
	"math/rand"
)

type RandomAlgorithm struct {
	myEtcdClient *common.MyEtcdClient
}

func GetRandomAlgorithm() *RandomAlgorithm {
	return &RandomAlgorithm{
		myEtcdClient: common.GetMyEtcdClient(),
	}
}

func (ra *RandomAlgorithm) GetPlayerId() (string, error) {
	resp := ra.myEtcdClient.Get(common.ETCD_PREFIX_PLAYER_INFO, clientv3.WithPrefix())
	players := *ra.myEtcdClient.ConvertToPlayerInfo(resp)
	if len(players) <= 0 {
		return "", errors.New("No players are alive")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return players[r.Intn(len(players))].Id, nil
}
