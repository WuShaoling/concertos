package etcd

import (
	"github.com/coreos/etcd/clientv3"
	"time"
	"sync"
	"log"
	"github.com/concertos/player/util"
	"encoding/json"
	"context"
	"os"
	"github.com/shirou/gopsutil/mem"
	"runtime"
	"github.com/concertos/common"
	"github.com/shortid"
)

type Player struct {
	client *clientv3.Client
	Info   *common.PlayerInfo
}

func GetSysInfo(info *common.PlayerInfo) {
	memory, _ := mem.VirtualMemory()
	info.Memory = memory.Total
	hostname, _ := os.Hostname()
	info.Hostname = hostname
	info.Ips = util.GetIps()
	info.Cpu = runtime.NumCPU()
}

func (p *Player) HeartBeat() {

	for {
		GetSysInfo(p.Info)
		log.Println(p.Info)
		key := common.ETCD_PREFIX_PLAYERS_INFO + p.Info.Id
		value, _ := json.Marshal(&p.Info)

		resp, err := p.client.Grant(context.TODO(), common.HEART_BEAT)
		if err != nil {
			log.Println(err)
		}
		_, err = p.client.Put(context.TODO(), key, string(value), clientv3.WithLease(resp.ID))
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second * common.HEART_BEAT)
	}
}

func NewPlayer() *Player {

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   common.GetEtcdPoints(),
		DialTimeout: 2 * time.Second,
	})

	if err != nil {
		log.Fatal("Error: new etcd client error:", err)
		return nil
	}
	//
	//conductor := &Conductor{
	//	client: *etcdClient,
	//}
	//return conductor
	//
	//
	//cfg := client.Config{
	//	Endpoints:               common.GetEtcdPoints(),
	//	Transport:               client.DefaultTransport,
	//	HeaderTimeoutPerRequest: time.Second,
	//}
	//
	//etcdClient, err := client.New(cfg)
	//if err != nil {
	//	log.Fatal("Error: cannot connec to etcd:", err)
	//}
	var info = new(common.PlayerInfo)
	info.Id = shortid.MustGenerate()
	info.State = common.ONLINE
	GetSysInfo(info)
	player := &Player{
		Info:   info,
		client: etcdClient,
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
