package conductor

import (
	"github.com/concertos/conductor/module/restapi"
	"github.com/concertos/conductor/module/scheduler"
	"github.com/concertos/conductor/module/manager"
	"sync"
	"github.com/concertos/conductor/module/websocket"
)

type Conductor struct {
	RestApi   *restapi.RestApi
	Rchedule  *scheduler.Scheduler
	Manager   *manager.Manager
	WebSocket *websocket.WebSocket
}

var once sync.Once
var conductor *Conductor

func GetConductor() *Conductor {
	once.Do(func() {
		conductor = &Conductor{
			RestApi:   restapi.GetRestApi(),
			Rchedule:  scheduler.GetScheduler(),
			Manager:   manager.GetManage(),
			WebSocket: websocket.GetWebSocket(),
		}
	})
	return conductor
}
