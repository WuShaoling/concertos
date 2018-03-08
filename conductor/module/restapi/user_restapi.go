package restapi

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"github.com/emicklei/go-restful-openapi"
	"github.com/concertos/common"
	"github.com/coreos/etcd/clientv3"
	"encoding/json"
	"time"
	"log"
)

type UserResource struct {
	myEtcdClient *common.MyEtcdClient
}

func GetUserResource() *UserResource {
	return &UserResource{
		myEtcdClient: common.GetMyEtcdClient(),
	}
}

func (u UserResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/users").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	tags := []string{"users"}

	ws.Route(ws.GET("/").To(u.getAllUsers).
		Doc("get all users").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]common.UserInfo{}).
		Returns(http.StatusOK, "OK", []common.UserInfo{}).
		Returns(http.StatusInternalServerError, "InternalServerError", nil))

	ws.Route(ws.GET("/{user-id}").To(u.getUser).
		Doc("get a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Writes(common.UserInfo{}).
		Returns(http.StatusOK, "OK", common.UserInfo{}).
		Returns(http.StatusInternalServerError, "InternalServerError", nil))

	ws.Route(ws.PUT("").To(u.createUser).
		Doc("create a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(common.UserInfo{}).
		Returns(http.StatusOK, "OK", nil).
		Returns(http.StatusInternalServerError, "InternalServerError", nil))

	return ws
}

func (ur *UserResource) getAllUsers(request *restful.Request, response *restful.Response) {
	resp, err := ur.myEtcdClient.Get(common.ETCD_PREFIX_USERS_INFO, clientv3.WithPrefix())
	if nil != err {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	log.Println(*ur.myEtcdClient.ConvertToUserInfo(resp))
	response.WriteEntity(*ur.myEtcdClient.ConvertToUserInfo(resp))
}

func (ur *UserResource) getUser(request *restful.Request, response *restful.Response) {
	resp, err := ur.myEtcdClient.Get(common.ETCD_PREFIX_USERS_INFO+request.PathParameter("user-id"), clientv3.WithPrefix())
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	log.Println(*ur.myEtcdClient.ConvertToUserInfo(resp))
	response.WriteEntity(*ur.myEtcdClient.ConvertToUserInfo(resp))
}

func (ur *UserResource) createUser(request *restful.Request, response *restful.Response) {
	//read content
	var user = new(common.UserInfo)
	err := request.ReadEntity(user)
	if err != nil { //read content error
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	user.Created = time.Now().Unix()

	//check is user-id exist
	resp, err := ur.myEtcdClient.Get(common.ETCD_PREFIX_USERS_INFO+request.PathParameter("user-id"), clientv3.WithPrefix())
	if nil != err {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	users := *ur.myEtcdClient.ConvertToUserInfo(resp)
	if len(users) > 0 {
		response.WriteErrorString(http.StatusFound, "User already exist")
		return
	}

	// put to etcd
	res, _ := json.Marshal(user)
	ur.myEtcdClient.Put(common.ETCD_PREFIX_USERS_INFO+user.Id, string(res), nil)

	response.WriteHeaderAndEntity(http.StatusOK, nil)
}
