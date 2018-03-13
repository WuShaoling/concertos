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
		Returns(http.StatusOK, "OK", []entity.PlayerInfo{}))

	ws.Route(ws.GET("/{playerid}").To(pr.getPlayer).
		Doc("get player according to player-id").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("playerid", "identifier of the player").DataType("string")).
		Writes([]entity.PlayerInfo{}).
		Returns(http.StatusOK, "OK", []entity.PlayerInfo{}))

	return ws
}

func (pm *PlayerResource) getAllPlayers(request *restful.Request, response *restful.Response) {
	players := *pm.myEtcdClient.ConvertToPlayerInfo(
		pm.myEtcdClient.Get(common.ETCD_PREFIX_PLAYER_INFO, clientv3.WithPrefix()))
	response.WriteEntity(players)
}

func (pm *PlayerResource) getPlayer(request *restful.Request, response *restful.Response) {
	key := common.ETCD_PREFIX_PLAYER_INFO + request.PathParameter("playerid")
	player := *pm.myEtcdClient.ConvertToPlayerInfo(pm.myEtcdClient.Get(key, clientv3.WithPrefix()))
	response.WriteEntity(player)
}

type PlayerResource struct {
	myEtcdClient *common.MyEtcdClient
}

func GetPlayerResource() *PlayerResource {
	return &PlayerResource{
		myEtcdClient: common.GetMyEtcdClient(),
	}
}
