package manager

import (
	"github.com/concertos/module/common"
)

//func getSysInfo(info *common.PlayerInfo) {
//	memory, _ := mem.VirtualMemory()
//	info.Memory = memory.Total
//
//	hostname, _ := os.Hostname()
//	info.Hostname = hostname
//
//	info.Ips = util.GetIps()
//	info.Cpu = runtime.NumCPU()
//}
//
//func (p *PlayerManager) KeepAlive() {
//	// set player info
//	key := common.ETCD_PREFIX_PLAYER_INFO + p.Info.Id
//	value, _ := json.Marshal(p.Info)
//	log.Println(key, string(value))
//	p.myEtcdClient.Put(key, string(value), nil)
//
//	interrupt := make(chan os.Signal, 1)
//	signal.Notify(interrupt, os.Interrupt)
//
//	ticker := time.NewTicker(time.Second * common.PLAYER_HEART_BEAT)
//	defer ticker.Stop()
//
//	for {
//		select {
//		case <-ticker.C:
//			key := common.ETCD_PREFIX_PLAYER_ALIVE + p.Info.Id
//			value, _ := json.Marshal(p.Info.Id)
//
//			log.Println(key, string(value))
//
//			if resp, err := p.myEtcdClient.GetClientV3().Grant(context.TODO(), common.PLAYER_HEART_BEAT); err != nil {
//				log.Println(err)
//			} else if _, err = p.myEtcdClient.GetClientV3().Put(context.TODO(), key, string(value), clientv3.WithLease(resp.ID)); nil != err {
//				log.Println(err)
//			}
//		case <-interrupt:
//			ticker.Stop()
//			log.Println("System interrupt, heart beat interrupt, ticker stop")
//			return
//		}
//	}
//}

type PlayerManager struct {
	myEtcdClient *common.MyEtcdClient
}

func GetPlayerManage() *PlayerManager {
	return &PlayerManager{
		myEtcdClient: common.GetMyEtcdClient(),
	}
}
