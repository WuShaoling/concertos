package websocket

import (
	"github.com/concertos/module/common"
	"encoding/json"
	"log"
	"github.com/concertos/player/util"
)

func (c *Client) HandleMsg(message []byte) {
	log.Println(string(message))
	var wsm = new(common.WebSocketMessage)
	json.Unmarshal(message, wsm)

	switch wsm.MessageType {
	case common.P_WS_REGISTER_PLAYER:
		c.Id = string(wsm.Content)
	case common.P_WS_REGISTER_USER:
		c.Id = string(wsm.Content)
	case common.P_WS_START_CONTAINER:
		c.startContainer(wsm)
	case common.P_WS_START_CONTAINER_RESULT:
	default:
		log.Println(wsm.MessageType, " ", wsm.Receiver, " ", wsm.Content)
	}
}

func (c *Client) startContainer(wsm *common.WebSocketMessage) {
	// get container-info from etcd
	etcd := common.GetMyEtcdClient()
	resp := etcd.Get(common.ETCD_PREFIX_CONTAINER_INFO + wsm.Content)
	container := *etcd.ConvertToContainerInfo(resp)

	// get player id, if player not alive, choose a new player
	playerid := container[0].PlayerId
	wsm.Content = string(util.MyJsonMarshal(container))

	// send start message to player
	ws := GetWebSocket()
	ws.WriteTo <- []byte(playerid)
	ws.WriteTo <- []byte(util.MyJsonMarshal(wsm))
}
