package etcd

import (
	"github.com/coreos/etcd/client"
	"github.com/concertos/common"
	"encoding/json"
	"log"
	"context"
)

func (c *Conductor) SetUser(info common.PlayerInfo) error {
	key := common.ETCD_PREFIX_USERS_INFO + info.Id
	value, _ := json.Marshal(info)
	_, err := c.KeysAPI.Set(context.Background(), key, string(value), &client.SetOptions{})
	if (err != nil) {
		log.Println("Error set player ", info.Id, " : ", err)
		return err
	}
	return nil
}

func (c *Conductor) DeleteUser(info common.PlayerInfo) error {
	return nil
}

func (c *Conductor) UpdateUser(info common.PlayerInfo) error {
	return nil
}
