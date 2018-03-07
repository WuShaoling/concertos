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
	ClientV3     *clientv3.Client
	MyEtcdClient *common.MyEtcdClient
	Info         *common.PlayerInfo
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
	// set player info
	key := common.ETCD_PREFIX_USERS_INFO + p.Info.Id
	value, _ := json.Marshal(p.Info)
	p.MyEtcdClient.Put(key, string(value), nil)

	// update player is alive
	for {
		key := common.ETCD_PREFIX_PLAYERS_ALIVE + p.Info.Id
		value, _ := json.Marshal(p.Info.Id)

		resp, err := p.ClientV3.Grant(context.TODO(), common.HEART_BEAT)
		if err != nil {
			log.Println(err)
		}
		_, err = p.MyEtcdClient.Put(key, string(value), clientv3.WithLease(resp.ID))
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second * common.HEART_BEAT)
	}
}

func NewPlayer() *Player {
	var info = new(common.PlayerInfo)
	info.Id = shortid.MustGenerate()
	info.State = common.ONLINE
	GetSysInfo(info)

	player := &Player{
		Info:         info,
		MyEtcdClient: common.GetMyEtcdClient(),
		ClientV3:     common.GetClientV3(),
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
