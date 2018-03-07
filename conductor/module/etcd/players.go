package etcd

import (
	"context"
	"encoding/json"
	"github.com/coreos/etcd/client"
	"log"
	"github.com/concertos/common"
)

func (c *Conductor) PlayerExpire(info *common.PlayerInfo) {
	info.State = common.PAUSED
	key := common.ETCD_PREFIX_PLAYERS_INFO + info.Id
	value, _ := json.Marshal(info)
	_, err := c.KeysAPI.Set(context.Background(), key, string(value), &client.SetOptions{})
	if err != nil {
		log.Println("Error PlayerExpire() ", info.Id, " : ", err)
	}
}

func (c *Conductor) SetPlayer(info *common.PlayerInfo) error {
	return nil
}

func (c *Conductor) DeletePlayer(info *common.PlayerInfo) error {
	return nil
}

func (c *Conductor) GetPlayer(id string) (*common.PlayerInfo, error) {
	return nil, nil
}

func (c *Conductor) Watch() {
	watcher := c.KeysAPI.Watcher(common.ETCD_PREFIX_PLAYERS_INFO, &client.WatcherOptions{
		Recursive: true,
	})
	for {
		res, err := watcher.Next(context.Background())
		if err != nil {
			log.Println("Error watch players:", err)
			break
		}
		switch  res.Action {
		case "expire":
			info := NodeToPlayerInfo(res.PrevNode)
			log.Println("Expire player ", info.Id)
			c.PlayerExpire(info)
			break;
		default:
			info := NodeToPlayerInfo(res.Node)
			log.Println("Watch player ", info.Id, " , ", res.Action)
			break;
		}
	}
}

func NodeToPlayerInfo(node *client.Node) *common.PlayerInfo {
	log.Println(node.Value)
	info := &common.PlayerInfo{}
	err := json.Unmarshal([]byte(node.Value), info)
	if err != nil {
		log.Println("Error NodeToPlayerInfo : ", err)
	}
	return info
}
