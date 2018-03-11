package restapi

import (
	"github.com/concertos/module/common"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"net/http"
	"github.com/coreos/etcd/clientv3"
	"github.com/concertos/module/entity"
)

func (pr *PlayerResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/players").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	tags := []string{"players"}

	ws.Route(ws.GET("/").To(pr.getAllPlayers).
		Doc("get all players").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.PlayerInfo{}).
		Returns(http.StatusOK, "OK", []entity.PlayerInfo{}).
		Returns(http.StatusInternalServerError, "InternalServerError", "error info"))

	ws.Route(ws.GET("/{playerid}").To(pr.getPlayer).
		Doc("get player according to playerid").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("playerid", "identifier of the player").DataType("string")).
		Writes([]entity.PlayerInfo{}).
		Returns(http.StatusOK, "OK", []entity.PlayerInfo{}).
		Returns(http.StatusInternalServerError, "InternalServerError", "error info"))

	ws.Route(ws.PUT("").To(pr.addPlayer).
		Doc("add a player").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.PlayerInfo{}).
		Returns(http.StatusOK, "OK", nil).
		Returns(http.StatusInternalServerError, "InternalServerError", "error info"))

	return ws
}

func (pm *PlayerResource) getAllPlayers(request *restful.Request, response *restful.Response) {
	if resp, err := pm.myEtcdClient.Get(common.ETCD_PREFIX_PLAYER_INFO, clientv3.WithPrefix()); nil != err {
		response.WriteError(http.StatusInternalServerError, err)
	} else {
		response.WriteEntity(*pm.myEtcdClient.ConvertToPlayerInfo(resp))
	}
}

func (pm *PlayerResource) getPlayer(request *restful.Request, response *restful.Response) {
	if resp, err := pm.myEtcdClient.Get(
		common.ETCD_PREFIX_PLAYER_INFO+request.PathParameter("playerid"), clientv3.WithPrefix()); nil != err {
		response.WriteError(http.StatusInternalServerError, err)
	} else {
		response.WriteEntity(*pm.myEtcdClient.ConvertToPlayerInfo(resp))
	}
}

func (pm *PlayerResource) addPlayer(request *restful.Request, response *restful.Response) {
	response.WriteEntity(nil)
}

type PlayerResource struct {
	myEtcdClient *common.MyEtcdClient
}

func GetPlayerResource() *PlayerResource {
	return &PlayerResource{
		myEtcdClient: common.GetMyEtcdClient(),
	}
}
