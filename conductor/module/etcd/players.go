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
	key := common.ETCD_PREFIX_PLAYERS_INFO + info.Id
	value, _ := json.Marshal(info)
	_, err := c.MyEtcdClent.Put(key, string(value), nil)
	if err != nil {
		log.Println("Error PutPlayer() ", info.Id, " : ", err)
	}
	return err
}

func (c *Conductor) DeletePlayer(id string) error {
	var p = new(common.PlayerInfo)
	p.Id = id
	p.State = common.OFFLINE
	key := common.ETCD_PREFIX_PLAYERS_INFO + id
	value, _ := json.Marshal(p)
	_, err := c.MyEtcdClent.Put(key, string(value), nil)
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
	rch := c.MyEtcdClent.GetClientV3().Watch(context.Background(), common.ETCD_PREFIX_PLAYERS_INFO, clientv3.WithPrefix())
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

//
//func (c *Conductor) PutPlayerId(id string) error {
//	_, err := c.client.Put(context.Background(), key, string(value), nil)
//	c.client.Delete()
//	if err != nil {
//		log.Println("Error PutPlayer() ", info.Id, " : ", err)
//	}
//	return err
//}
