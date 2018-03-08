package restapi

import (
	"log"
	"net/http"
	"github.com/emicklei/go-restful"
	"sync"
)

type RestApi struct {
	playerResource    *PlayerResource
	userResource      *UserResource
	containerResource *ContainerResource
}

func (rest *RestApi) Start() {
	user := UserResource{}
	restful.DefaultContainer.Add(user.WebService())

	log.Printf("start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var restApi *RestApi
var once sync.Once

func GetRestApi() *RestApi {
	once.Do(func() {
		restApi = &RestApi{
			userResource:      GetUserResource(),
			playerResource:    GetPlayerResource(),
			containerResource: GetContainerResource(),
		}
	})
	return restApi
}
