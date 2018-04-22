package websocket

import (
	"github.com/concertos/module/common"
	"encoding/json"
	"log"
	"github.com/concertos/player/module/executor"
	"github.com/concertos/module/entity"
	"github.com/concertos/player/module/manager"
	"github.com/concertos/module/util"
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
	case common.P_WS_CONTAINER_INSTALL:
		wsm.Content = executor.Install(ws.getContainer([]byte(wsm.Content)))
		wsm.MessageType = common.P_WS_CONTAINER_INSTALL_R

	case common.P_WS_CONTAINER_START:
		wsm.Content = executor.Start(ws.getContainer([]byte(wsm.Content)))
		wsm.MessageType = common.P_WS_CONTAINER_START_R

	case common.P_WS_CONTAINER_STOP:
		wsm.Content = executor.Stop(ws.getContainer([]byte(wsm.Content)))
		wsm.MessageType = common.P_WS_CONTAINER_STOP_R

	case common.P_WS_CONTAINER_REMOVE:
		log.Println("-------------------------------", wsm.Content)
		wsm.Content = executor.Remove(ws.getContainer([]byte(wsm.Content)))
		wsm.MessageType = common.P_WS_CONTAINER_REMOVE_R

	default:
		log.Println("Unknown message type")
	}

	wsm.Receiver = wsm.Sender
	wsm.Sender = manager.GetManage().PlayerManager.Info.Id
	ws.Send <- util.MyJsonMarshal(wsm)
}
