package websocket

import (
	"github.com/concertos/module/common"
	"github.com/concertos/module/util"
	"github.com/concertos/conductor/module/scheduler"
)

func (c *Client) startContainer(wsm *common.WebSocketMessage) {
	// get container-info from etcd
	etcd := common.GetMyEtcdClient()
	resp := etcd.Get(common.ETCD_PREFIX_CONTAINER_INFO + wsm.Content)
	containers := *etcd.ConvertToContainerInfo(resp)

	// get player id
	// if player is not available, call schedule module to select a new player
	playerid := containers[0].PlayerId
	if etcd.CheckExist(common.ETCD_PREFIX_PLAYER_ALIVE+playerid) == false {
		containers[0].PlayerId = scheduler.GetScheduler().RandomAlgorithm.GetPlayerId()
	}
	wsm.Content = string(util.MyJsonMarshal(containers[0]))

	// send start message to player
	ws := GetWebSocket()
	ws.WriteTo <- []byte(containers[0].PlayerId)
	ws.WriteTo <- []byte(util.MyJsonMarshal(wsm))
}
