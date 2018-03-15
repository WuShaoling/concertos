package websocket

import (
	"github.com/concertos/module/common"
	"encoding/json"
	"github.com/concertos/player/util"
	"log"
	"github.com/concertos/module/entity"
	"github.com/concertos/player/module/executor"
)

func (ws *WebSocket) HandleMsg(message []byte) {

	log.Println(string(message))
	var wsm = new(common.WebSocketMessage)
	json.Unmarshal(message, wsm)

	switch wsm.MessageType {
	case common.P_WS_INSTALL_CONTAINER:
		ws.startContainer(wsm)
	case common.P_WS_REGISTER_PLAYER:
	case common.P_WS_START_CONTAINER:
	case common.P_WS_START_CONTAINER_RESULT:

	}
}

func (ws *WebSocket) startContainer(wsm *common.WebSocketMessage) {
	// Get container
	container := new(entity.ContainerInfo)
	json.Unmarshal([]byte(wsm.Content), container)

	// call executor module
	resp := executor.GetExecutor().RegExecutor.Start(container)

	log.Println(resp)

	// Response result
	wsm.Content = resp
	wsm.Receiver = wsm.Sender
	wsm.Sender = container.PlayerId
	wsm.MessageType = common.P_WS_START_CONTAINER_RESULT
	GetWebSocket().Send <- util.MyJsonMarshal(wsm)
}
