package websocket

import (
	"github.com/concertos/module/common"
	"encoding/json"
	"log"
)

func (c *Client) HandleMsg(message []byte) {
	log.Println("HandleMsg : ", string(message))
	var wsm = new(common.WebSocketMessage)
	json.Unmarshal(message, wsm)

	switch wsm.MessageType {
	case common.P_WS_REGISTER_PLAYER:
		c.Id = string(wsm.Sender)
		log.Println(c.Id)

	case common.P_WS_REGISTER_USER:
		c.Id = string(wsm.Sender)
		log.Println(c.Id)

	case common.P_WS_INSTALL_CONTAINER:
		c.installContainer(wsm)

	case common.P_WS_INSTALL_CONTAINER_R:
		ws := GetWebSocket()
		ws.WriteTo <- []byte(wsm.Receiver)
		ws.WriteTo <- message

	case common.P_WS_START_CONTAINER:
		c.startContainer(wsm)

	case common.P_WS_STOP_CONTAINER:
		c.startContainer(wsm)

	case common.P_WS_REMOVE_CONTAINER:
		c.startContainer(wsm)

	default:
		log.Println(wsm.MessageType, " ", wsm.Receiver, " ", wsm.Content)
	}
}
