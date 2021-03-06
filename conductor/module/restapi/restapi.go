package restapi

import (
	"log"
	"net/http"
	"github.com/emicklei/go-restful"
	"sync"
	"github.com/concertos/module/common"
)

func (rest *RestApi) Start() {

	pr := PlayerResource{}
	restful.DefaultContainer.Add(pr.WebService())

	cr := ContainerResource{}
	restful.DefaultContainer.Add(cr.WebService())

	sr := StaticResource{}
	restful.DefaultContainer.Add(sr.WebService())

	log.Printf("rest api server listening on: " + common.RESTAPI_ADDR)
	log.Fatal(http.ListenAndServe(common.RESTAPI_ADDR, nil))
}

var restApi *RestApi
var once sync.Once

type RestApi struct {
	PlayerResource    *PlayerResource
	ContainerResource *ContainerResource
	StaticResource    *StaticResource
}

func GetRestApi() *RestApi {
	once.Do(func() {
		restApi = &RestApi{
			PlayerResource:    GetPlayerResource(),
			ContainerResource: GetContainerResource(),
			StaticResource:    GetStaticResource(),
		}
	})
	return restApi
}
