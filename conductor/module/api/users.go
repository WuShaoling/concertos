package api

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"github.com/emicklei/go-restful-openapi"
	"context"
	"log"
	"encoding/json"
	"github.com/ventu-io/go-shortid"
	"strings"
	"github.com/concertos/common"
	"github.com/concertos/conductor/module/etcd"
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

	ws.Route(ws.PUT("/{user-id}").To(u.updateUser).
		Doc("update a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Reads(common.UserInfo{}))

	ws.Route(ws.PUT("").To(u.createUser).
		Doc("create a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(common.UserInfo{}))

	ws.Route(ws.DELETE("/{user-id}").To(u.removeUser).
		Doc("delete a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")))

	return ws
}

func (u UserResource) findAllUsers(request *restful.Request, response *restful.Response) {
	c := etcd.GetConductor()
	resp, err := c.KeysAPI.Get(context.Background(), "/users", nil)
	if err != nil {
		log.Println("err read users")
	}
	var users []common.UserInfo
	for _, v := range resp.Node.Nodes {
		var user common.UserInfo
		json.Unmarshal([]byte(v.Value), &user)
		arr := strings.Split(v.Key, "/")
		user.Id = arr[len(arr)-1]
		users = append(users, user)
	}
	response.WriteEntity(users)
}

func (u UserResource) findUser(request *restful.Request, response *restful.Response) {
	c := etcd.GetConductor()
	id := request.PathParameter("user-id")
	resp, err := c.KeysAPI.Get(context.Background(), "/users/"+id, nil)
	if err != nil {
		log.Println("cannot get user from etcd")
		response.WriteErrorString(http.StatusNotFound, "UserInfo could not be found.")
		return
	}
	var user common.UserInfo
	json.Unmarshal([]byte(resp.Node.Value), &user)
	response.WriteEntity(user)
}

func (u *UserResource) updateUser(request *restful.Request, response *restful.Response) {
	c := etcd.GetConductor()
	usr := new(common.UserInfo)
	err := request.ReadEntity(&usr)
	if err == nil {
		str, _ := json.Marshal(usr)
		c.KeysAPI.Set(context.Background(), "/users/"+usr.Id, string(str), nil)
		response.WriteHeaderAndEntity(http.StatusCreated, usr)
		response.WriteEntity(usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func (u *UserResource) createUser(request *restful.Request, response *restful.Response) {
	c := etcd.GetConductor()
	sid, _ := shortid.New(1, shortid.DefaultABC, 2342)
	uid, _ := sid.Generate()
	usr := common.UserInfo{Id: uid}
	err := request.ReadEntity(&usr)
	if err == nil {
		str, _ := json.Marshal(usr)
		c.KeysAPI.Set(context.Background(), "/users/"+usr.Id, string(str), nil)
		response.WriteHeaderAndEntity(http.StatusCreated, usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func (u *UserResource) removeUser(request *restful.Request, response *restful.Response) {
	c := etcd.GetConductor()
	id := request.PathParameter("user-id")

	c.KeysAPI.Delete(context.Background(), "/users/"+id, nil)
}
