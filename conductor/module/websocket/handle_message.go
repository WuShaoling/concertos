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

	case common.P_WS_CONTAINER_INSTALL:
		c.installContainer(wsm)
	case common.P_WS_CONTAINER_INSTALL_R:
		log.Println("handle install container result")
		GetWebSocket().HttpWait[wsm.Receiver] <- wsm.Content
		delete(GetWebSocket().HttpWait, wsm.Receiver)

	case common.P_WS_CONTAINER_START:
		c.startContainer(wsm)
	case common.P_WS_CONTAINER_START_R:
		log.Println("handle start container result")
		GetWebSocket().HttpWait[wsm.Receiver] <- wsm.Content
		delete(GetWebSocket().HttpWait, wsm.Receiver)

	case common.P_WS_CONTAINER_STOP:
		c.startContainer(wsm)
	case common.P_WS_CONTAINER_STOP_R:
		log.Println("handle stop container result")
		GetWebSocket().HttpWait[wsm.Receiver] <- wsm.Content
		delete(GetWebSocket().HttpWait, wsm.Receiver)

	case common.P_WS_CONTAINER_REMOVE:
		c.startContainer(wsm)
	case common.P_WS_CONTAINER_REMOVE_R:
		log.Println("handle remove container result")
		GetWebSocket().HttpWait[wsm.Receiver] <- wsm.Content
		delete(GetWebSocket().HttpWait, wsm.Receiver)

	default:
		log.Println("HandleMsg : ", string(message))
	}
}
