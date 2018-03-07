package etcd

import (
	"context"
	"encoding/json"
	"log"
	"github.com/concertos/common"
	"github.com/coreos/etcd/clientv3"
	"fmt"
)

func (c *Conductor) PutPlayer(info *common.PlayerInfo) error {
	return nil
}

func (c *Conductor) DeletePlayer(id string) error {
	var p = new(common.PlayerInfo)
	p.Id = id
	p.State = common.OFFLINE
	key := common.ETCD_PREFIX_PLAYERS_INFO + id
	value, _ := json.Marshal(p)
	_, err := c.client.Put(context.Background(), key, string(value), nil)
	if err != nil {
		log.Println("Error PlayerExpire() ", p.Id, " : ", err)
	}
	return nil
}

func (c *Conductor) GetPlayer(id string) (*common.PlayerInfo, error) {
	return nil, nil
}

func (c *Conductor) GetAllPlayer() ([]common.PlayerInfo, error) {
	return nil, nil
}

func (c *Conductor) Watch() {
	rch := c.client.Watch(context.Background(), common.ETCD_PREFIX_PLAYERS_INFO, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch  ev.Type.String() {
			case "DELETE":
				c.DeletePlayer(string(ev.Kv.Key))
				break;
			default:
				break;
			}
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

//func NodeToPlayerInfo(node []byte) *common.PlayerInfo {
//	info := &common.PlayerInfo{}
//	err := json.Unmarshal(node, info)
//	if err != nil {
//		log.Println("Error NodeToPlayerInfo : ", err)
//	}
//	return info
//}
