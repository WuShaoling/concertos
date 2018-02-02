package etcd

import (
	"context"
	"encoding/json"
	"github.com/coreos/etcd/client"
	"log"
	"sync"
	"time"
	"github.com/concertos/config"
	"github.com/concertos/player/module/etcd"
)

var conductor *Conductor
var once sync.Once

type Conductor struct {
	Players map[string]*Player
	KeysAPI client.KeysAPI
}

type Player struct {
	Online bool
	Info   etcd.ClientInfo
}

func NodeToPlayerInfo(node *client.Node) *etcd.ClientInfo {
	log.Println(node.Value)
	info := &etcd.ClientInfo{}
	err := json.Unmarshal([]byte(node.Value), info)
	if err != nil {
		log.Print(err)
	}
	return info
}

func (c *Conductor) UpdatePlayer(info *etcd.ClientInfo, online bool) {
	player, ok := c.Players[info.Id]
	if ok {
		player.Online = online
	}
}

func (c *Conductor) AddPlayer(info etcd.ClientInfo) {
	player := &Player{
		Online: true,
		Info:   info,
	}
	c.Players[player.Info.Id] = player
}

func (c *Conductor) DeletePlayer(info *etcd.ClientInfo) {
	delete(c.Players, info.Id)
}

func (c *Conductor) Watch() {
	api := c.KeysAPI
	watcher := api.Watcher("/players/", &client.WatcherOptions{
		Recursive: true,
	})
	for {
		res, err := watcher.Next(context.Background())
		if err != nil {
			log.Println("Error watch workers:", err)
			break
		}
		if res.Action == "expire" {
			info := NodeToPlayerInfo(res.PrevNode)
			log.Println("Expire player ", info.Id)
			c.UpdatePlayer(info, false)
		} else if res.Action == "set" {
			info := NodeToPlayerInfo(res.Node)
			if _, ok := c.Players[info.Id]; ok {
				log.Println("Update player ", info.Id)
				c.UpdatePlayer(info, true)
			} else {
				log.Println("Add player ", info.Id)
				c.AddPlayer(*info)
			}
		} else if res.Action == "delete" {
			info := NodeToPlayerInfo(res.Node)
			log.Println("Delete player ", info.Id)
			c.DeletePlayer(info)
		} else if res.Action == "get" {
			log.Fatal("res.Action == get")
		} else if res.Action == "update" {
			log.Fatal("res.Action == update")
		} else if res.Action == "create" {
			log.Fatal("res.Action == create")
		} else if res.Action == "compareAndSwap" {
			log.Fatal("res.Action == compareAndSwap")
		} else if res.Action == "compareAndDelete" {
			log.Fatal("res.Action == compareAndDelete")
		}
	}
}

func NewConductor() *Conductor {
	cfg := client.Config{
		Endpoints:               config.GetEtcdPoints(),
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := client.New(cfg)
	if err != nil {
		log.Fatal("Error: cannot connec to etcd:", err)
	}

	conductor := &Conductor{
		Players: make(map[string]*Player),
		KeysAPI: client.NewKeysAPI(etcdClient),
	}
	return conductor
}

func GetConductor() *Conductor {
	once.Do(func() {
		conductor = NewConductor()
	})
	return conductor
}
