package websocket

import (
	"github.com/concertos/module/common"
	"github.com/concertos/player/util"
)

func (c *Client) startContainer(wsm *common.WebSocketMessage) {
	// get container-info from etcd, wsm.Content is container-id
	etcd := common.GetMyEtcdClient()
	resp := etcd.Get(common.ETCD_PREFIX_CONTAINER_INFO + wsm.Content)
	container := *etcd.ConvertToContainerInfo(resp)

	// get player id, if player not alive, choose a new player
	playerid := container[0].PlayerId
	// if player not available, call schedule module to select a new player
	// func () {}

	wsm.Content = string(util.MyJsonMarshal(container))

	// send start message to player
	ws := GetWebSocket()
	ws.WriteTo <- []byte(playerid)
	ws.WriteTo <- []byte(util.MyJsonMarshal(wsm))
}
