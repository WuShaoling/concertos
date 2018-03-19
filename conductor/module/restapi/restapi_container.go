package restapi

import (
	"github.com/concertos/module/common"
	"github.com/emicklei/go-restful"
	"net/http"
)

func (cr *ContainerResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/containers").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.POST("/install").To(cr.installContainer))

	ws.Route(ws.POST("/stop").To(cr.stopContainer))

	ws.Route(ws.POST("/remove").To(cr.stopContainer))

	return ws
}

func (cr *ContainerResource) installContainer(request *restful.Request, response *restful.Response) {
	//// 1. get container from request
	//var container = new(entity.ContainerInfo)
	//if err := request.ReadEntity(container); err != nil {
	//	response.WriteError(http.StatusInternalServerError, err)
	//	return
	//}
	//
	//// 2. set other info
	//container.Id = shortid.MustGenerate()
	//container.Ip = dccp.GetDccp().GetIp()
	//container.State = common.CONTAINER_STATE_STOPPED
	//container.Created = time.Now().Unix()
	//container.PlayerId = scheduler.GetScheduler().RandomAlgorithm.GetPlayerId()
	//
	//// 3. put to etcd
	//cr.myEtcdClient.Put(common.ETCD_PREFIX_CONTAINER_INFO+container.Id,
	//	string(util.MyJsonMarshal(container)), nil)
	//
	//// 4 return result
	response.WriteHeaderAndEntity(http.StatusOK, nil)
}

func (cr *ContainerResource) stopContainer(req *restful.Request, response *restful.Response) {
	response.WriteHeaderAndEntity(http.StatusOK, nil)
}

func (cr *ContainerResource) removeContainer(req *restful.Request, response *restful.Response) {
	response.WriteHeaderAndEntity(http.StatusOK, nil)
}

type ContainerResource struct {
	myEtcdClient *common.MyEtcdClient
}

func GetContainerResource() *ContainerResource {
	return &ContainerResource{
		myEtcdClient: common.GetMyEtcdClient(),
	}
}
