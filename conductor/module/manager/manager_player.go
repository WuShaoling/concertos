package manager

import (
	"context"
	"encoding/json"
	"log"
	"github.com/concertos/module/common"
	"github.com/coreos/etcd/clientv3"
	"errors"
	"strings"
)

func (pm *PlayerManager) expirePlayer(id string) error {
	// 1. get player info accounding to alive id
	strs := strings.Split(id, "/")
	resp, err := pm.myEctdClient.Get(common.ETCD_PREFIX_PLAYER_INFO + strs[len(strs)-1])
	if nil != err {
		return err
	}
	players := *pm.myEctdClient.ConvertToPlayerInfo(resp)
	if len(players) <= 0 {
		return errors.New("Error : player " + id + " not exist")
	}

	// 2. update state
	players[0].State = common.PLAYER_STATE_OFFLINE
	if player, err := json.Marshal(players[0].State); err != nil {
		return err
	} else if _, err := pm.myEctdClient.Put(
		common.ETCD_PREFIX_PLAYER_INFO+players[0].Id, string(player), nil); nil != err {
		return err
	}
	return nil
}

func (pm *PlayerManager) Watch() {
	rch := pm.myEctdClient.GetClientV3().Watch(context.Background(), common.ETCD_PREFIX_PLAYER_ALIVE, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			log.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			switch  ev.Type.String() {
			case "DELETE":
				log.Println(pm.expirePlayer(string(ev.Kv.Key)))
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
