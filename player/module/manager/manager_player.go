package manager

import (
	"github.com/concertos/module/common"
	"github.com/shirou/gopsutil/mem"
	"os"
	"github.com/concertos/player/util"
	"encoding/json"
	"github.com/coreos/etcd/clientv3"
	"time"
	"runtime"
	"log"
	"context"
	"github.com/shortid"
)

func getSysInfo(info *common.PlayerInfo) {
	memory, _ := mem.VirtualMemory()
	info.Memory = memory.Total

	hostname, _ := os.Hostname()
	info.Hostname = hostname

	info.Ips = util.GetIps()
	info.Cpu = runtime.NumCPU()
}

func (p *PlayerManager) HeartBeat() {
	// set player info
	key := common.ETCD_PREFIX_PLAYER_INFO + p.Info.Id
	value, _ := json.Marshal(p.Info)
	log.Println(key, value)
	p.myEtcdClient.Put(key, string(value), nil)

	// update player is alive
	for {
		key := common.ETCD_PREFIX_PLAYER_ALIVE + p.Info.Id
		value, _ := json.Marshal(p.Info.Id)

		log.Println(key, value)

		resp, err := p.myEtcdClient.GetClientV3().Grant(context.TODO(), common.HEART_BEAT)
		if err != nil {
			log.Println(err)
		}
		_, err = p.myEtcdClient.Put(key, string(value), clientv3.WithLease(resp.ID))
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second * common.HEART_BEAT)
	}
}

type PlayerManager struct {
	myEtcdClient *common.MyEtcdClient
	Info         *common.PlayerInfo
}

func GetPlayerManage() *PlayerManager {
	var info = new(common.PlayerInfo)
	info.Id = shortid.MustGenerate()
	info.State = common.PLAYER_STATE_ONLINE
	getSysInfo(info)

	return &PlayerManager{
		Info:         info,
		myEtcdClient: common.GetMyEtcdClient(),
	}
}
