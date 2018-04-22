package restapi

import (
	"github.com/emicklei/go-restful"
	"path"
	"fmt"
	"net/http"
	"log"
)

var rootdir = "./ui/"

func (sr *StaticResource) WebService() *restful.WebService {

	ws := new(restful.WebService)

	ws.Route(ws.GET("/static/{subpath:*}").To(sr.staticFromPathParam))
	ws.Route(ws.GET("/static").To(sr.staticFromQueryParam))

	return ws
}

func (sr *StaticResource) staticFromPathParam(req *restful.Request, resp *restful.Response) {
	actual := path.Join(rootdir, req.PathParameter("subpath"))
	fmt.Printf("serving %s ... (from %s)\n", actual, req.PathParameter("subpath"))
	http.ServeFile(resp.ResponseWriter, req.Request, actual)
}

func (sr *StaticResource) staticFromQueryParam(req *restful.Request, resp *restful.Response) {
	log.Println(req.QueryParameter("resource"))
	http.ServeFile(resp.ResponseWriter, req.Request, path.Join(rootdir, req.QueryParameter("resource")))
}

type StaticResource struct {
}

func GetStaticResource() *StaticResource {
	return &StaticResource{
	}
}
