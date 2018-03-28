package websocket

import (
	"github.com/concertos/module/common"
	"github.com/concertos/player/util"
	"github.com/concertos/conductor/module/scheduler"
)

func (c *Client) startContainer(wsm *common.WebSocketMessage) {
	// get container-info from etcd
	etcd := common.GetMyEtcdClient()
	resp := etcd.Get(common.ETCD_PREFIX_CONTAINER_INFO + wsm.Content)
	container := *etcd.ConvertToContainerInfo(resp)

	// get player id
	// if player is not available, call schedule module to select a new player
	playerid := container[0].PlayerId
	if etcd.CheckExist(common.ETCD_PREFIX_PLAYER_ALIVE+playerid) == false {
		container[0].PlayerId = scheduler.GetScheduler().RandomAlgorithm.GetPlayerId()
	}
	wsm.Content = string(util.MyJsonMarshal(container[0]))

	// send start message to player
	ws := GetWebSocket()
	ws.WriteTo <- []byte(playerid)
	ws.WriteTo <- []byte(util.MyJsonMarshal(wsm))
}
