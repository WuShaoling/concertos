package websocket

import (
	"github.com/concertos/module/common"
	"github.com/concertos/module/util"
)

func (c *Client) stopContainer(wsm *common.WebSocketMessage) {

	//根据id获得容器信息
	resp := c.myEtcdClient.Get(common.ETCD_PREFIX_CONTAINER_INFO + wsm.Content)
	container := (*c.myEtcdClient.ConvertToContainerInfo(resp))[0]

	wsm.MessageType = common.P_WS_CONTAINER_STOP
	wsm.Receiver = container.PlayerId
	wsm.Content = string(util.MyJsonMarshal(container))

	ws := GetWebSocket()
	ws.WriteTo <- []byte(wsm.Receiver)
	ws.WriteTo <- util.MyJsonMarshal(wsm)

}
