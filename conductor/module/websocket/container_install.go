package websocket

import (
	"github.com/concertos/module/entity"
	"github.com/concertos/module/common"
	"github.com/concertos/conductor/module/scheduler"
	"github.com/concertos/module/util"
	"github.com/shortid"
	"encoding/json"
)

func (c *Client) installContainer(wsm *common.WebSocketMessage) {

	//Name      string
	//Command   string
	//Describe  string
	//BaseImage string
	//CPU       int
	//Memory    uint64

	//Id        string
	//PlayerId  string
	//State     int

	// 启动时设置
	//Ip        string

	ws := GetWebSocket()

	wsm.MessageType = common.P_WS_CONTAINER_INSTALL_R
	wsm.Receiver = wsm.Sender
	wsm.Sender = ""

	// get and set info
	var container = new(entity.ContainerInfo)
	json.Unmarshal([]byte(wsm.Content), container)

	// check if name exist
	if c.myEtcdClient.CheckExist(common.ETCD_PREFIX_CONTAINER_INFO + container.Name) {
		wsm.Content = "Container " + container.Name + "already exist!"
		ws.WriteTo <- []byte(wsm.Receiver)
		ws.WriteTo <- util.MyJsonMarshal(wsm)
		return
	}

	// set other info and put to etcd
	container.Id = shortid.MustGenerate()
	container.State = common.CONTAINER_STATE_STOPPED
	container.PlayerId = scheduler.GetScheduler().RandomAlgorithm.GetPlayerId()
	c.myEtcdClient.Put(common.ETCD_PREFIX_CONTAINER_INFO+container.Name, string(util.MyJsonMarshal(*container)))

	wsm.Content = "ok"

	ws.WriteTo <- []byte(wsm.Receiver)
	ws.WriteTo <- util.MyJsonMarshal(wsm)
}
