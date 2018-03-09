package player

import (
	"sync"
	"github.com/concertos/player/module/manager"
	"github.com/concertos/player/module/websocket"
)

type Player struct {
	Manager   *manager.Manager
	WebSocket *websocket.WebSocket
}

var player *Player
var once sync.Once

func GetEtcdClient() *Player {
	once.Do(func() {
		player = &Player{
			Manager:   manager.GetManage(),
			WebSocket: websocket.GetWetSocket(),
		}
	})
	return player
}
