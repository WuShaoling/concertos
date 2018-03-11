package restapi

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"github.com/emicklei/go-restful-openapi"
	"github.com/concertos/module/common"
	"github.com/coreos/etcd/clientv3"
	"encoding/json"
	"time"
	"github.com/shortid"
	"github.com/concertos/module/entity"
)

func (u UserResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/users").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	tags := []string{"users"}

	ws.Route(ws.GET("/").To(u.getAllUsers).
		Doc("get all users").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.UserInfo{}).
		Returns(http.StatusOK, "OK", []entity.UserInfo{}).
		Returns(http.StatusInternalServerError, "InternalServerError", "error info"))

	ws.Route(ws.GET("/{userid}").To(u.getUser).
		Doc("get a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("userid", "identifier of the user").DataType("string")).
		Writes(entity.UserInfo{}).
		Returns(http.StatusOK, "OK", entity.UserInfo{}).
		Returns(http.StatusInternalServerError, "InternalServerError", "error info"))

	ws.Route(ws.PUT("").To(u.createUser).
		Doc("create a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.UserInfo{}).
		Returns(http.StatusOK, "OK", nil).
		Returns(http.StatusFound, "User already exist", "error info").
		Returns(http.StatusInternalServerError, "InternalServerError", "error info"))

	return ws
}

func (ur *UserResource) getAllUsers(request *restful.Request, response *restful.Response) {
	if resp, err := ur.myEtcdClient.Get(common.ETCD_PREFIX_USER_INFO, clientv3.WithPrefix()); nil != err {
		response.WriteError(http.StatusInternalServerError, err)
	} else {
		response.WriteEntity(*ur.myEtcdClient.ConvertToUserInfo(resp))
	}
}

func (ur *UserResource) getUser(request *restful.Request, response *restful.Response) {
	if resp, err := ur.myEtcdClient.Get(common.ETCD_PREFIX_USER_INFO+request.PathParameter("userid"), nil); err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	} else {
		response.WriteEntity(*ur.myEtcdClient.ConvertToUserInfo(resp))
	}
}

func (ur *UserResource) createUser(request *restful.Request, response *restful.Response) {
	// 1. read content
	var user = new(entity.UserInfo)
	if err := request.ReadEntity(user); err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	user.Id = shortid.MustGenerate()
	user.Created = time.Now().Unix()

	// 2. put to etcd
	res, _ := json.Marshal(user)
	if _, err := ur.myEtcdClient.Put(common.ETCD_PREFIX_USER_INFO+user.Id, string(res), nil); err != nil {
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	} else {
		response.WriteHeaderAndEntity(http.StatusOK, nil)
	}
}

type UserResource struct {
	myEtcdClient *common.MyEtcdClient
}

func GetUserResource() *UserResource {
	return &UserResource{
		myEtcdClient: common.GetMyEtcdClient(),
	}
}
