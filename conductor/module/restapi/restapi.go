package restapi

import (
	"log"
	"net/http"
	"github.com/emicklei/go-restful"
	"sync"
)

func (rest *RestApi) Start() {
	ur := UserResource{}
	restful.DefaultContainer.Add(ur.WebService())

	pr := PlayerResource{}
	restful.DefaultContainer.Add(pr.WebService())

	cr := ContainerResource{}
	restful.DefaultContainer.Add(cr.WebService())

	sr := StaticResource{}
	restful.DefaultContainer.Add(sr.WebService())

	log.Printf("Start rest api, listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var restApi *RestApi
var once sync.Once

type RestApi struct {
	PlayerResource    *PlayerResource
	UserResource      *UserResource
	ContainerResource *ContainerResource
	StaticResource    *StaticResource
}

func GetRestApi() *RestApi {
	once.Do(func() {
		restApi = &RestApi{
			UserResource:      GetUserResource(),
			PlayerResource:    GetPlayerResource(),
			ContainerResource: GetContainerResource(),
			StaticResource:    GetStaticResource(),
		}
	})
	return restApi
}
