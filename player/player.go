package player

import (
	"sync"
	"github.com/concertos/player/module/manager"
	"github.com/concertos/player/module/websocket"
	"github.com/concertos/module/common"
	"log"
	"github.com/concertos/module/util"
	"github.com/concertos/player/module/executor"
)

type Player struct {
	myEtcdClient *common.MyEtcdClient
	Manager      *manager.Manager
	WebSocket    *websocket.WebSocket
	Executor     *executor.Executor
}

func (p *Player) Register() {
	// send register msg to conductor
	wsm := common.WebSocketMessage{
		MessageType: common.P_WS_REGISTER_PLAYER,
		Content:     "",
		Receiver:    "",
		Sender:      p.Manager.PlayerManager.Info.Id,
	}
	json := util.MyJsonMarshal(wsm)
	p.WebSocket.Send <- json
	log.Println("Send register msg to conductor: ", wsm)

	// write msg to etcd
	value := string(util.MyJsonMarshal(p.Manager.PlayerManager.Info))
	p.myEtcdClient.Put(common.ETCD_PREFIX_PLAYER_INFO+p.Manager.PlayerManager.Info.Id, value)
	log.Println("Put palyer info to etcd: ", value)
}

var player *Player
var once sync.Once

func GetPlayer() *Player {
	once.Do(func() {
		player = &Player{
			Manager:      manager.GetManage(),
			WebSocket:    websocket.GetWebSocket(),
			Executor:     executor.GetExecutor(),
			myEtcdClient: common.GetMyEtcdClient(),
		}
	})
	return player
}
