package etcd

import (
	"github.com/coreos/etcd/client"
	"time"
	"sync"
	"log"
	"github.com/concertos/player/util"
	"encoding/json"
	"context"
	"os"
	"github.com/shirou/gopsutil/mem"
	"github.com/shortid"
	"runtime"
	"github.com/concertos/common"
)

type Player struct {
	KeysAPI client.KeysAPI
	Info    common.PlayerInfo
}

func GetSysInfo() common.PlayerInfo {
	hostname, _ := os.Hostname()
	memory, _ := mem.VirtualMemory()
	id := shortid.MustGenerate()
	ips := util.GetIps()
	cpu := runtime.NumCPU()

	return common.PlayerInfo{
		Id:       id,
		Ips:      ips,
		Hostname: hostname,
		Memory:   memory.Total,
		Cpu:      cpu,
	}
}

func (p *Player) HeartBeat() {
	for {
		p.Info = GetSysInfo()
		key := common.ETCD_PREFIX_PLAYERS_INFO + p.Info.Id
		value, _ := json.Marshal(&p.Info)

		_, err := p.KeysAPI.Set(context.Background(), key, string(value), &client.SetOptions{
			TTL: time.Second * common.TTL,
		})
		if err != nil {
			log.Println("Error update EtcdClientInfo:", err)
		}
		time.Sleep(time.Second * common.HEART_BEAT)
	}
}

func NewPlayer() *Player {
	cfg := client.Config{
		Endpoints:               common.GetEtcdPoints(),
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := client.New(cfg)
	if err != nil {
		log.Fatal("Error: cannot connec to etcd:", err)
	}

	player := &Player{
		Info:    GetSysInfo(),
		KeysAPI: client.NewKeysAPI(etcdClient),
	}
	return player
}

var player *Player
var once sync.Once

func GetEtcdClient() *Player {
	once.Do(func() {
		player = NewPlayer()
	})
	return player
}
