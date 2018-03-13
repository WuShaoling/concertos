package websocket

import (
	"github.com/concertos/module/common"
	"encoding/json"
	"github.com/concertos/player/util"
	"log"
)

//type WebSocketMessage struct {
//	MessageType int
//	Receiver    string
//	Content     string
//}

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
	// get container-info from etcd
	etcd := common.GetMyEtcdClient()
	resp := etcd.Get(common.ETCD_PREFIX_CONTAINER_INFO + wsm.Content)
	container := *etcd.ConvertToContainerInfo(resp)

	// get player id, if player not alive, choose a new player
	playerid := container[0].PlayerId
	wsm.Content = string(util.MyJsonMarshal(container))
	log.Println(playerid)

	//// send start message to player
	//ws := GetWebSocket()
	//ws.WriteTo <- []byte(playerid)
	//ws.WriteTo <- []byte(util.MyJsonMarshal(wsm))
}
