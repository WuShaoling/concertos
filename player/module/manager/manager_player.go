package manager

import (
	"github.com/concertos/module/common"
	"os"
	"os/signal"
	"time"
	"github.com/concertos/module/util"
	"github.com/coreos/etcd/clientv3"
	"log"
	"context"
	"github.com/concertos/module/entity"
	"github.com/shirou/gopsutil/mem"
	"io/ioutil"
	"github.com/shortid"
	"runtime"
)

func setSysInfo() *entity.PlayerInfo {

	info := new(entity.PlayerInfo)

	id, err := ioutil.ReadFile(".config/player-id")
	if nil != err || len(string(id)) == 0 {
		log.Println(err)
		id1 := shortid.MustGenerate()
		info.Id = id1
		util.WriteFile(id1, ".config/player-id")
	} else {
		info.Id = string(id)
	}

	memory, _ := mem.VirtualMemory()
	info.Memory = memory.Total

	hostname, _ := os.Hostname()
	info.Hostname = hostname

	info.Ips = util.GetIps()
	info.Cpu = runtime.NumCPU()

	return info
}

func (pm *PlayerManager) KeepAlive() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ticker := time.NewTicker(time.Second * common.PLAYER_HEART_BEAT)
	defer ticker.Stop()

	key := common.ETCD_PREFIX_PLAYER_ALIVE + pm.Info.Id
	value := string(util.MyJsonMarshal(pm.Info.Id))

	for {
		select {
		case <-ticker.C:
			resp, err := pm.myEtcdClient.GetClientV3().Grant(context.TODO(), common.PLAYER_TTL)
			if err != nil {
				log.Println(err)
			} else if _, err = pm.myEtcdClient.GetClientV3().Put(context.TODO(),
				key, value, clientv3.WithLease(resp.ID)); nil != err {
				log.Println(err)
			}
		case <-interrupt:
			log.Println("System interrupt, keep alive interrupt, ticker stop")
			return
		}
	}
}

type PlayerManager struct {
	myEtcdClient *common.MyEtcdClient
	Info         *entity.PlayerInfo
}

func GetPlayerManager() *PlayerManager {
	return &PlayerManager{
		myEtcdClient: common.GetMyEtcdClient(),
		Info:         setSysInfo(),
	}
}
