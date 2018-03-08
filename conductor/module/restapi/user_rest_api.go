package restapi

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"github.com/emicklei/go-restful-openapi"
	"github.com/concertos/common"
	"log"
	"github.com/coreos/etcd/clientv3"
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
		Returns(200, "OK", []common.UserInfo{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/{user-id}").To(u.getUser).
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

func (u *UserResource) getAllUsers(request *restful.Request, response *restful.Response) {
	//var pi *interface{}
	pi := new(interface{})
	*pi = &[]common.PlayerInfo{}

	err := u.myEtcdClient.Get(pi, common.ETCD_PREFIX_PLAYERS_INFO, clientv3.WithPrefix())
	log.Println(*pi)
	log.Println(err)

	//c := etcd.GetConductor()
	//user, err := c.GetUser(request.PathParameter("user-id"))
	//if err != nil {
	//	response.WriteError(http.StatusInternalServerError, err)
	//	return
	//}
	response.WriteEntity(nil)
	//var users []common.PlayerInfo
	//err := u.myEtcdClient.Get(&users, common.ETCD_PREFIX_PLAYERS_INFO, clientv3.WithPrefix())
	//if nil != err {
	//	log.Println(err)
	//	response.WriteError(http.StatusInternalServerError, err)
	//}
	//log.Println(users)
	//c := etcd.GetConductor()
	//users, err := c.GetAllUser()
	//if err != nil {
	//	response.WriteError(http.StatusInternalServerError, err)
	//	return
	//}
	//response.WriteEntity(users)
}

func (u *UserResource) getUser(request *restful.Request, response *restful.Response) {
	var users []common.PlayerInfo
	var a *interface{}
	a = new(interface{})
	*a = users

	err := u.myEtcdClient.Get(a, common.ETCD_PREFIX_PLAYERS_INFO, clientv3.WithPrefix())
	log.Println(users)
	log.Println(err)

	//c := etcd.GetConductor()
	//user, err := c.GetUser(request.PathParameter("user-id"))
	//if err != nil {
	//	response.WriteError(http.StatusInternalServerError, err)
	//	return
	//}
	print(users)
	response.WriteEntity(nil)
}

func (u *UserResource) createUser(request *restful.Request, response *restful.Response) {
	//c := etcd.GetConductor()
	//var user = new(common.UserInfo)
	//err := request.ReadEntity(user)
	//if err != nil { //read content error
	//	log.Println("Error createUser() : ", err)
	//	response.WriteError(http.StatusInternalServerError, err)
	//	return
	//}
	//
	//u1, err1 := c.GetUser(user.Id)
	//
	//if err1 != nil { // InternalServerError
	//	if c.GetErrorType(err) != 100 {
	//		log.Println("Error StatusInternalServerError : ", err1)
	//		response.WriteError(http.StatusInternalServerError, err1)
	//		return
	//	}
	//} else if u1.Id == user.Id { //user id already exist
	//	log.Println("Error user id already exist")
	//	response.WriteErrorString(http.StatusOK, "User already exist")
	//	return
	//}
	//
	//err2 := c.SetUser(user)
	//if nil != err2 { // set user error
	//	log.Println("Error set user : ", err2)
	//	response.WriteError(http.StatusInternalServerError, err2)
	//	return
	//}

	response.WriteHeaderAndEntity(http.StatusCreated, nil)
}
