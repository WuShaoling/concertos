package restapi

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"github.com/emicklei/go-restful-openapi"
	"github.com/concertos/common"
	"github.com/concertos/conductor/module/etcd"
	"log"
	"github.com/concertos/util"
)

type UserResource struct {
}

func (u UserResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/users").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	tags := []string{"users"}

	ws.Route(ws.GET("/").To(u.findAllUsers).
		Doc("get all users").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]common.UserInfo{}).
		Returns(200, "OK", []common.UserInfo{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/{user-id}").To(u.findUser).
		Doc("get a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Writes(common.UserInfo{}).
		Returns(200, "OK", common.UserInfo{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("").To(u.createUser).
		Doc("create a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(common.UserInfo{}))

	return ws
}

func (u *UserResource) findAllUsers(request *restful.Request, response *restful.Response) {
	c := etcd.GetConductor()
	users, err := c.GetAllUser()
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	response.WriteEntity(users)
}

func (u *UserResource) findUser(request *restful.Request, response *restful.Response) {
	c := etcd.GetConductor()
	user, err := c.GetUser(request.PathParameter("user-id"))
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	response.WriteEntity(user)
}

func (u *UserResource) createUser(request *restful.Request, response *restful.Response) {
	c := etcd.GetConductor()
	var user = new(common.UserInfo)
	err := request.ReadEntity(user)
	if err != nil { //read content error
		log.Println("Error createUser() : ", err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	u1, err1 := c.GetUser(user.Id)

	if err1 != nil { // InternalServerError
		if util.GetEtcdErrorType(err1.Error())[0] != "100" {
			log.Println("Error StatusInternalServerError : ", err1)
			response.WriteError(http.StatusInternalServerError, err1)
			return
		}
	} else if u1.Id == user.Id { //user id already exist
		log.Println("Error user id already exist")
		response.WriteErrorString(http.StatusOK, "User already exist")
		return
	}

	err2 := c.SetUser(user)
	if nil != err2 { // set user error
		log.Println("Error set user : ", err2)
		response.WriteError(http.StatusInternalServerError, err2)
		return
	}

	response.WriteHeaderAndEntity(http.StatusCreated, user)
}
