package websocket

import (
	"github.com/concertos/module/common"
	"encoding/json"
	"log"
	"github.com/concertos/player/module/executor"
	"github.com/concertos/module/entity"
	"github.com/concertos/player/module/manager"
	"github.com/concertos/player/util"
)

func (ws *WebSocket) getContainer(data []byte) *entity.ContainerInfo {
	container := new(entity.ContainerInfo)
	json.Unmarshal(data, container)
	return container
}

func (ws *WebSocket) HandleMsg(message []byte) {
	log.Println("HandleMsg : ", string(message))
	var wsm = new(common.WebSocketMessage)
	json.Unmarshal(message, wsm)

	executor := executor.GetExecutor().DockerExecutor

	switch wsm.MessageType {
	case common.P_WS_INSTALL_CONTAINER:
		wsm.Content = executor.Install(ws.getContainer([]byte(wsm.Content)))
		wsm.Receiver = wsm.Sender
		wsm.Sender = manager.GetManage().PlayerManager.Info.Id
		wsm.MessageType = common.P_WS_INSTALL_CONTAINER_R
		ws.Send <- util.MyJsonMarshal(wsm)

	case common.P_WS_START_CONTAINER:
		executor.Start(ws.getContainer([]byte(wsm.Content)))

	case common.P_WS_STOP_CONTAINER:
		executor.Stop(ws.getContainer([]byte(wsm.Content)))

	case common.P_WS_REMOVE_CONTAINER:
		executor.Remove(ws.getContainer([]byte(wsm.Content)))

	default:
		log.Println("Unknown message type")
	}
}
