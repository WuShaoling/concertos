package websocket

import (
	"github.com/concertos/module/entity"
	"github.com/concertos/module/dccp"
	"github.com/concertos/module/common"
	"time"
	"github.com/concertos/conductor/module/scheduler"
	"github.com/concertos/player/util"
	"github.com/shortid"
	"encoding/json"
	"log"
)

//type ContainerInfo struct {
//	Name      string `json:"Name" description:"name of container'"`
//	User      string `json:"User" description:"user id of container"`
//	Describe  string `json:"Describe" description:"additional description information"`
//	BaseImage string `json:"BaseImage description:""`
//	Command   string `json:"Command" description:"command"`

//	Id        string `json:"Id" description:"uniquely identifies of container"`
//	Ip        string `json:"Ip" description:"container's ip"`
//	State     int    `json:"State" description:"the current status of the container, running, stopped, paused..."`
//	Created   int64  `json:"Created" description:"created time"`
//	PlayerId  string `json:"PlayerId" description:""`
//}

func (c *Client) installContainer(wsm *common.WebSocketMessage) {
	// get and set info
	var container = new(entity.ContainerInfo)
	json.Unmarshal([]byte(wsm.Content), container)

	log.Println(container)

	container.Id = shortid.MustGenerate()
	container.Ip = dccp.GetDccp().GetIp()
	container.State = common.CONTAINER_STATE_STOPPED
	container.Created = time.Now().Unix()
	container.PlayerId = scheduler.GetScheduler().RandomAlgorithm.GetPlayerId()

	log.Println("installContainer : ", string(util.MyJsonMarshal(*container)))

	// put to etcd
	c.myEtcdClient.Put(common.ETCD_PREFIX_CONTAINER_INFO+container.Id, string(util.MyJsonMarshal(*container)))

	// build message and call player to pull image if image not exist
	wsm.Receiver = container.PlayerId
	wsm.Content = string(util.MyJsonMarshal(container))
	ws := GetWebSocket()
	ws.WriteTo <- []byte(container.PlayerId)
	ws.WriteTo <- util.MyJsonMarshal(wsm)
}
