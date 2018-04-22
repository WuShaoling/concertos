package restapi

import (
	"github.com/concertos/module/common"
	"github.com/emicklei/go-restful"
	"github.com/coreos/etcd/clientv3"
	"strings"
	"github.com/concertos/module/entity"
	"strconv"
	"log"
)

func (pr *PlayerResource) WebService() *restful.WebService {

	ws := new(restful.WebService)
	ws.Path("/players").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/").To(pr.getAllPlayers))
	ws.Route(ws.GET("/getall").To(pr.getAllPlayersContainers))
	ws.Route(ws.GET("/getindex").To(pr.getIndexPageData))
	ws.Route(ws.GET("/getbyid/{playerid}").To(pr.getPlayer))

	return ws
}

func (pm *PlayerResource) getIndexPageData(request *restful.Request, response *restful.Response) {

	result := make(map[string]string)
	var total_cpu int
	var total_mem uint64

	var used_cpu int
	var used_mem uint64

	total_cpu = 0
	total_mem = 0
	used_cpu = 0
	used_mem = 0

	players := *pm.myEtcdClient.ConvertToPlayerInfo(
		pm.myEtcdClient.Get(common.ETCD_PREFIX_PLAYER_INFO, clientv3.WithPrefix()))
	for _, pv := range players {
		total_cpu += pv.Cpu
		total_mem += pv.Memory
	}

	containers := *pm.myEtcdClient.ConvertToContainerInfo(
		pm.myEtcdClient.Get(common.ETCD_PREFIX_CONTAINER_INFO, clientv3.WithPrefix()))
	for _, pv := range containers {
		used_cpu += pv.CPU
		used_mem += pv.Memory
	}

	result["total_cpu"] = strconv.Itoa(total_cpu)
	result["total_memory"] = strconv.FormatUint(total_mem, 10)
	result["total_player"] = strconv.Itoa(len(players))
	result["total_container"] = strconv.Itoa(len(players))
	result["used_cpu"] = strconv.Itoa(used_cpu)
	result["used_memory"] = strconv.FormatUint(used_mem, 10)

	log.Println(result)

	response.WriteEntity(result)
}

func (pm *PlayerResource) getAllPlayers(request *restful.Request, response *restful.Response) {
	players := *pm.myEtcdClient.ConvertToPlayerInfo(
		pm.myEtcdClient.Get(common.ETCD_PREFIX_PLAYER_INFO, clientv3.WithPrefix()))
	resp_state := pm.myEtcdClient.Get(common.ETCD_PREFIX_PLAYER_ALIVE, clientv3.WithPrefix())

	for index, pv := range players {
		players[index].Status = 0
		for _, v := range resp_state.Kvs {
			if strings.Compare(string(v.Value), "\""+pv.Id+"\"") == 0 {
				players[index].Status = 1
				break
			}
		}
	}
	response.WriteEntity(players)
}

func (pm *PlayerResource) getAllPlayersContainers(request *restful.Request, response *restful.Response) {
	contaienrs := *pm.myEtcdClient.ConvertToContainerInfo(
		pm.myEtcdClient.Get(common.ETCD_PREFIX_CONTAINER_INFO, clientv3.WithPrefix()))

	resp_states := pm.myEtcdClient.Get(common.ETCD_PREFIX_CONTAINER_RUNNING, clientv3.WithPrefix())

	for index, pv := range contaienrs {
		contaienrs[index].State = 0
		for _, v := range resp_states.Kvs {
			if strings.Compare(string(v.Value), "\""+pv.Id+"\"") == 0 {
				contaienrs[index].State = 1
				break
			}
		}
	}

	var cons = make(map[string][]entity.ContainerInfo)
	for _, pv := range contaienrs {
		cons[pv.PlayerId] = append(cons[pv.PlayerId], pv)
	}

	response.WriteEntity(cons)
}

func (pm *PlayerResource) getPlayer(request *restful.Request, response *restful.Response) {
	playerid := request.PathParameter("playerid")
	player := *pm.myEtcdClient.ConvertToPlayerInfo(pm.myEtcdClient.Get(common.ETCD_PREFIX_PLAYER_INFO + playerid))
	resp_state := pm.myEtcdClient.Get(common.ETCD_PREFIX_PLAYER_ALIVE + playerid)

	if len(player) == 0 {
		response.WriteEntity(player)
		return
	} else {
		player[0].Status = 0
		if resp_state.Count != 0 {
			player[0].Status = 1
		}
	}

	response.WriteEntity(player[0])
}

type PlayerResource struct {
	myEtcdClient *common.MyEtcdClient
}

func GetPlayerResource() *PlayerResource {
	return &PlayerResource{
		myEtcdClient: common.GetMyEtcdClient(),
	}
}
