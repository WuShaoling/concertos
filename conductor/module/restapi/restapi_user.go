package restapi

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"github.com/emicklei/go-restful-openapi"
	"github.com/concertos/module/common"
	"github.com/coreos/etcd/clientv3"
	"time"
	"github.com/shortid"
	"github.com/concertos/module/entity"
	"github.com/concertos/player/util"
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
	users := *ur.myEtcdClient.ConvertToUserInfo(
		ur.myEtcdClient.Get(common.ETCD_PREFIX_USER_INFO, clientv3.WithPrefix()))
	response.WriteEntity(users)
}

func (ur *UserResource) getUser(request *restful.Request, response *restful.Response) {
	key := common.ETCD_PREFIX_USER_INFO + request.PathParameter("userid")
	user := *ur.myEtcdClient.ConvertToUserInfo(ur.myEtcdClient.Get(key))
	response.WriteEntity(user)
}

func (ur *UserResource) createUser(request *restful.Request, response *restful.Response) {
	// 1. read content
	var user = new(entity.UserInfo)
	request.ReadEntity(user)

	// 2. set other info
	user.Id = shortid.MustGenerate()
	user.Created = time.Now().Unix()

	// 3. put to etcd
	ur.myEtcdClient.Put(common.ETCD_PREFIX_USER_INFO+user.Id, string(util.MyJsonMarshal(user)))

	// 4. return
	response.WriteHeaderAndEntity(http.StatusOK, nil)
}

type UserResource struct {
	myEtcdClient *common.MyEtcdClient
}

func GetUserResource() *UserResource {
	return &UserResource{
		myEtcdClient: common.GetMyEtcdClient(),
	}
}
