package manager

import (
	"context"
	"log"
	"github.com/concertos/module/common"
	"github.com/coreos/etcd/clientv3"
	"strings"
	"github.com/concertos/player/util"
)

func (pm *PlayerManager) expirePlayer(id string) {
	log.Println("Expire Player " + id)

	strs := strings.Split(id, "/")
	resp := pm.myEctdClient.Get(common.ETCD_PREFIX_PLAYER_INFO + strs[len(strs)-1])
	players := *pm.myEctdClient.ConvertToPlayerInfo(resp)

	players[0].State = common.PLAYER_STATE_OFFLINE

	pm.myEctdClient.Put(common.ETCD_PREFIX_PLAYER_INFO+players[0].Id, string(util.MyJsonMarshal(players[0])))
}

func (pm *PlayerManager) Watch() {
	rch := pm.myEctdClient.GetClientV3().Watch(context.Background(), common.ETCD_PREFIX_PLAYER_ALIVE, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			//log.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			switch  ev.Type.String() {
			case "DELETE":
				pm.expirePlayer(string(ev.Kv.Key))
			default:
			}
		}
	}
}

type PlayerManager struct {
	myEctdClient *common.MyEtcdClient
}

func GetPlayerManager() *PlayerManager {
	return &PlayerManager{}
}
