package websocket

import (
	"github.com/concertos/module/common"
	"github.com/concertos/module/util"
)

func (c *Client) removeContainer(wsm *common.WebSocketMessage) {

	//根据id获得容器信息
	resp := c.myEtcdClient.Get(common.ETCD_PREFIX_CONTAINER_INFO + wsm.Content)
	containers := *c.myEtcdClient.ConvertToContainerInfo(resp)
	container := containers[0]

	wsm.MessageType = common.P_WS_CONTAINER_REMOVE
	wsm.Receiver = container.PlayerId
	wsm.Content = string(util.MyJsonMarshal(container))

	ws := GetWebSocket()
	ws.WriteTo <- []byte(wsm.Receiver)
	ws.WriteTo <- util.MyJsonMarshal(wsm)

}
