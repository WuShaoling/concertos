package conductor

import (
	"context"
	"encoding/json"
	"github.com/coreos/etcd/client"
	"log"
	"sync"
	"time"
)

const ETCDENDPOINT = "http://127.0.0.1:2379"

var conductor *Conductor
var once sync.Once

type Conductor struct {
	Players map[string]*Player
	KeysAPI client.KeysAPI
}

type Player struct {
	Online bool
	Info   PlayerInfo
}

type PlayerInfo struct {
	Id       string
	Ips      []string
	Hostname string
	Memory   uint64
	Cpu      int
}

func NodeToPlayerInfo(node *client.Node) *PlayerInfo {
	log.Println(node.Value)
	info := &PlayerInfo{}
	err := json.Unmarshal([]byte(node.Value), info)
	if err != nil {
		log.Print(err)
	}
	return info
}

func (c *Conductor) UpdatePlayer(info *PlayerInfo, online bool) {
	player, ok := c.Players[info.Id]
	if ok {
		player.Online = online
	}
}

func (c *Conductor) AddPlayer(info PlayerInfo) {
	player := &Player{
		Online: true,
		Info:   info,
	}
	c.Players[player.Info.Id] = player
}

func (c *Conductor) DeletePlayer(info *PlayerInfo) {
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
		Endpoints:               []string{ETCDENDPOINT},
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
