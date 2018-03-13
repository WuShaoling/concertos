package player

import (
	"sync"
	"github.com/concertos/player/module/manager"
	"github.com/concertos/player/module/websocket"
	"github.com/concertos/module/common"
	"encoding/json"
	"os"
	"os/signal"
	"time"
	"github.com/coreos/etcd/clientv3"
	"log"
	"github.com/shortid"
	"github.com/shirou/gopsutil/mem"
	"github.com/concertos/player/util"
	"runtime"
	"context"
	"github.com/concertos/module/entity"
)

type Player struct {
	myEtcdClient *common.MyEtcdClient
	Info         *entity.PlayerInfo
	Manager      *manager.Manager
	WebSocket    *websocket.WebSocket
}

func (p *Player) Register() {
	// send register msg to conductor
	msg, _ := json.Marshal(&common.WebSocketMessage{
		MessageType: common.P_WS_REGISTER_PLAYER,
		Content:     p.Info.Id,
		Receiver:    "",
	})
	p.WebSocket.Send <- msg

	// write msg to etcd
	key := common.ETCD_PREFIX_PLAYER_INFO + p.Info.Id
	value, _ := json.Marshal(p.Info)
	log.Println(key, string(value))
	p.myEtcdClient.Put(key, string(value), nil)
}

func getSysInfo(info *entity.PlayerInfo) *entity.PlayerInfo {
	info.Id = shortid.MustGenerate()
	info.State = common.PLAYER_STATE_ONLINE

	memory, _ := mem.VirtualMemory()
	info.Memory = memory.Total

	hostname, _ := os.Hostname()
	info.Hostname = hostname

	info.Ips = util.GetIps()
	info.Cpu = runtime.NumCPU()

	return info
}

func (p *Player) KeepAlive() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ticker := time.NewTicker(time.Second * common.HEART_BEAT)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			key := common.ETCD_PREFIX_PLAYER_ALIVE + p.Info.Id
			value, _ := json.Marshal(p.Info.Id)

			log.Println(key, string(value))
			resp, err := p.myEtcdClient.GetClientV3().Grant(context.TODO(), common.TTL)
			if err != nil {
				log.Println(err)
			} else if _, err = p.myEtcdClient.GetClientV3().Put(context.TODO(), key, string(value), clientv3.WithLease(resp.ID)); nil != err {
				log.Println(err)
			}
		case <-interrupt:
			log.Println("System interrupt, heart beat interrupt, ticker stop")
			return
		}
	}
}

var player *Player
var once sync.Once

func GetPlayer() *Player {
	once.Do(func() {
		player = &Player{
			Manager:      manager.GetManage(),
			WebSocket:    websocket.GetWebSocket(),
			Info:         getSysInfo(new(entity.PlayerInfo)),
			myEtcdClient: common.GetMyEtcdClient(),
		}
	})
	return player
}
