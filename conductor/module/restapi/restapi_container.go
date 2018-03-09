package restapi

import (
	"github.com/concertos/module/common"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"net/http"
	"github.com/ventu-io/go-shortid"
	"time"
	"github.com/concertos/conductor/module/scheduler"
	"errors"
	"github.com/concertos/module/dccp"
	"encoding/json"
)

func (cr *ContainerResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/containers").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	tags := []string{"containers"}

	ws.Route(ws.POST("/install").To(cr.installContainer).
		Doc("install a container").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(common.ContainerInfo{}).
		Returns(http.StatusOK, "OK", nil).
		Returns(http.StatusInternalServerError, "InternalServerError", "error info"))

	ws.Route(ws.POST("/start/").To(cr.startContainer).
		Doc("start a container").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("player-id", "identifier of the player").DataType("string")).
		Returns(http.StatusOK, "OK", nil).
		Returns(http.StatusInternalServerError, "InternalServerError", "error info"))

	ws.Route(ws.POST("/stop").To(cr.stopContainer).
		Doc("stop a container").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(common.UserInfo{}).
		Returns(http.StatusOK, "OK", nil).
		Returns(http.StatusInternalServerError, "InternalServerError", "error info"))

	ws.Route(ws.POST("/removeContainer").To(cr.stopContainer).
		Doc("remove a container").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(common.UserInfo{}).
		Returns(http.StatusOK, "OK", nil).
		Returns(http.StatusInternalServerError, "InternalServerError", "error info"))

	return ws
}

func (cr *ContainerResource) installContainer(request *restful.Request, response *restful.Response) {
	// 1. get container from request
	var container = new(common.ContainerInfo)
	if err := request.ReadEntity(container); err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	// 2. set other info
	container.Id = shortid.MustGenerate()
	if ip, err := dccp.GetDccp().GetIp(); nil != err {
		response.WriteError(http.StatusInternalServerError, errors.New("no ip resource available"))
		return
	} else {
		container.Ip = ip
	}
	container.State = common.CONTAINER_STATE_STOPPED
	container.Created = time.Now().Unix()
	if playerid, err := scheduler.GetScheduler().RandomAlgorithm.GetPlayerId(); nil != err {
		response.WriteError(http.StatusInternalServerError, err)
		return
	} else {
		container.PlayerId = playerid
	}

	// 3. put to etcd
	key := common.ETCD_PREFIX_CONTAINER_INFO + container.User + "/" + container.PlayerId + "/" + container.Id
	value, _ := json.Marshal(container)
	if _, err := cr.myEtcdClient.Put(key, string(value), nil); err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, nil)
}

// player manage module waill receive put msg, then start

func (cr *ContainerResource) startContainer(req *restful.Request, resp *restful.Response) {
	// 1. put to etcd waitting start field


	//



}

func (cr *ContainerResource) stopContainer(req *restful.Request, resp *restful.Response) {

}

func (cr *ContainerResource) removeContainer(req *restful.Request, resp *restful.Response) {

}

type ContainerResource struct {
	myEtcdClient *common.MyEtcdClient
}

func GetContainerResource() *ContainerResource {
	return &ContainerResource{
		myEtcdClient: common.GetMyEtcdClient(),
	}
}
