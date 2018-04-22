package restapi

import (
	"github.com/concertos/module/common"
	"github.com/emicklei/go-restful"
	"net/http"
	"github.com/coreos/etcd/clientv3"
	"strings"
	"github.com/concertos/module/entity"
	"github.com/concertos/module/util"
	"github.com/shortid"
	"github.com/concertos/conductor/module/scheduler"
	"github.com/concertos/conductor/module/websocket"
	"time"
	"log"
)

func (cr *ContainerResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/containers").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/").To(cr.getAll))

	ws.Route(ws.PUT("/install").To(cr.installContainer))
	ws.Route(ws.GET("/start/{container-name}").To(cr.startContainer))
	ws.Route(ws.GET("/stop/{container-name}").To(cr.stopContainer))
	ws.Route(ws.GET("/remove/{container-name}").To(cr.removeContainer))
	ws.Route(ws.GET("/cloudware/{cloudware-name}").To(cr.getCloudware))

	return ws
}

func (cr *ContainerResource) getCloudware(req *restful.Request, response *restful.Response) {
	response.WriteEntity("")
}

func (cr *ContainerResource) getAll(req *restful.Request, response *restful.Response) {
	contaienrs := *cr.myEtcdClient.ConvertToContainerInfo(
		cr.myEtcdClient.Get(common.ETCD_PREFIX_CONTAINER_INFO, clientv3.WithPrefix()))

	resp_states := cr.myEtcdClient.Get(common.ETCD_PREFIX_CONTAINER_RUNNING, clientv3.WithPrefix())

	for index, pv := range contaienrs {
		contaienrs[index].State = 0
		for _, v := range resp_states.Kvs {
			if strings.Compare(string(v.Value), pv.Name) == 0 {
				contaienrs[index].State = 1
				break
			}
		}
	}

	response.WriteEntity(contaienrs)
}

func (cr *ContainerResource) installContainer(request *restful.Request, response *restful.Response) {
	container := new(entity.ContainerInfo)
	if err := request.ReadEntity(container); err == nil {
		// check if name exist
		if cr.myEtcdClient.CheckExist(common.ETCD_PREFIX_CONTAINER_INFO + container.Name) {
			response.Write([]byte("Container name " + container.Name + "already exist!"))
			return
		}
		// set other info and put to etcd
		container.Id = shortid.MustGenerate()
		container.State = common.CONTAINER_STATE_STOPPED
		container.PlayerId = scheduler.GetScheduler().RandomAlgorithm.GetPlayerId()
		cr.myEtcdClient.Put(common.ETCD_PREFIX_CONTAINER_INFO+container.Name, string(util.MyJsonMarshal(*container)))

		response.WriteEntity(http.StatusOK)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func (cr *ContainerResource) startContainer(request *restful.Request, response *restful.Response) {
	//get params
	cname := request.PathParameter("container-name")

	log.Println("cname : " + cname)

	// get container-info from etcd
	etcd := common.GetMyEtcdClient()
	resp := etcd.Get(common.ETCD_PREFIX_CONTAINER_INFO + cname)
	containers := *etcd.ConvertToContainerInfo(resp)
	if len(containers) == 0 {
		log.Println("no such container")
		response.WriteEntity("no such container")
		return
	}
	container := containers[0]

	// check players
	if len(websocket.GetWebSocket().Clients) == 0 {
		log.Println("no player available")
		response.WriteEntity("no player available")
		return
	}

	// get player id
	// if player is not available, call schedule module to select a new player
	if etcd.CheckExist(common.ETCD_PREFIX_PLAYER_ALIVE+container.PlayerId) == false {
		container.PlayerId = scheduler.GetScheduler().RandomAlgorithm.GetPlayerId()
	}

	// construact websocket message
	wsm := common.WebSocketMessage{
		Receiver:    container.PlayerId,
		Sender:      string(time.Now().Unix()),
		MessageType: common.P_WS_CONTAINER_START,
		Content:     string(util.MyJsonMarshal(container)),
	}

	// send start message to player
	ws := websocket.GetWebSocket()
	c := make(chan string);
	ws.AddWait <- websocket.HttpWaitMsg{
		Id:      wsm.Sender,
		Channel: c,
	}
	ws.WriteTo <- []byte(wsm.Receiver)
	ws.WriteTo <- []byte(util.MyJsonMarshal(wsm))

	// wait for result
	result := <-c
	response.Write([]byte(result))
}

func (cr *ContainerResource) stopContainer(request *restful.Request, response *restful.Response) {
	//get params
	cname := request.PathParameter("container-name")

	// get container-info from etcd
	etcd := common.GetMyEtcdClient()
	resp := etcd.Get(common.ETCD_PREFIX_CONTAINER_INFO + cname)
	containers := *etcd.ConvertToContainerInfo(resp)
	if len(containers) == 0 {
		response.WriteEntity("no such container")
		return
	}
	container := containers[0]

	ponline := false;
	for k, _ := range websocket.GetWebSocket().Clients {
		if k.Id == container.PlayerId {
			ponline = true;
			break
		}
	}
	if !ponline {
		response.WriteEntity("error, player offline")
		return
	}

	// construact websocket message
	wsm := common.WebSocketMessage{
		Receiver:    container.PlayerId,
		Sender:      "" + string(time.Now().Unix()),
		MessageType: common.P_WS_CONTAINER_STOP,
		Content:     string(util.MyJsonMarshal(container)),
	}

	// send start message to player
	ws := websocket.GetWebSocket()
	c := make(chan string);
	ws.AddWait <- websocket.HttpWaitMsg{
		Id:      wsm.Sender,
		Channel: c,
	}
	ws.WriteTo <- []byte(wsm.Receiver)
	ws.WriteTo <- []byte(util.MyJsonMarshal(wsm))

	// wait for result
	result := <-c
	response.Write([]byte(result))
}

func (cr *ContainerResource) removeContainer(request *restful.Request, response *restful.Response) {
	//get params
	cname := request.PathParameter("container-name")

	// get container-info from etcd
	etcd := common.GetMyEtcdClient()
	resp := etcd.Get(common.ETCD_PREFIX_CONTAINER_INFO + cname)
	containers := *etcd.ConvertToContainerInfo(resp)
	if len(containers) == 0 {
		response.WriteEntity("no such container")
		return
	}
	container := containers[0]

	// if player offline, return error
	ponline := false
	for k, _ := range websocket.GetWebSocket().Clients {
		if k.Id == container.PlayerId {
			ponline = true
			break
		}
	}
	log.Println(ponline)
	if !ponline {
		response.WriteEntity("error, player offline")
		return
	}

	// construact websocket message
	wsm := common.WebSocketMessage{
		Receiver:    container.PlayerId,
		Sender:      string(time.Now().Unix()),
		MessageType: common.P_WS_CONTAINER_REMOVE,
		Content:     string(util.MyJsonMarshal(container)),
	}

	// send start message to player
	ws := websocket.GetWebSocket()
	c := make(chan string);
	ws.AddWait <- websocket.HttpWaitMsg{
		Id:      wsm.Sender,
		Channel: c,
	}
	ws.WriteTo <- []byte(wsm.Receiver)
	ws.WriteTo <- []byte(util.MyJsonMarshal(wsm))

	// wait for result
	result := <-c
	response.Write([]byte(result))
}

type ContainerResource struct {
	myEtcdClient *common.MyEtcdClient
}

func GetContainerResource() *ContainerResource {
	return &ContainerResource{
		myEtcdClient: common.GetMyEtcdClient(),
	}
}
