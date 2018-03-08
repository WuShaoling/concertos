package restapi

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"github.com/emicklei/go-restful-openapi"
	"github.com/concertos/common"
	"log"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/error"
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

func (ur *UserResource) getAllUsers(request *restful.Request, response *restful.Response) {
	users := new(interface{})
	*users = &[]common.PlayerInfo{}
	err := ur.myEtcdClient.Get(users, common.ETCD_PREFIX_PLAYERS_INFO, clientv3.WithPrefix())
	if err != nil {
		log.Println(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	response.WriteEntity(*users)
}

func (ur *UserResource) getUser(request *restful.Request, response *restful.Response) {
	users := new(interface{})
	*users = &[]common.PlayerInfo{}
	err := ur.myEtcdClient.Get(users,
		common.ETCD_PREFIX_PLAYERS_INFO+request.PathParameter("user-id"), clientv3.WithPrefix())
	if err != nil {
		log.Println(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	response.WriteEntity(*users)
}

func (ur *UserResource) createUser(request *restful.Request, response *restful.Response) {

	//read content
	var user = new(common.UserInfo)
	err := request.ReadEntity(user)
	if err != nil { //read content error
		log.Println(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	// check if userid exist
	users := new(interface{})
	*users = &[]common.PlayerInfo{}
	err = ur.myEtcdClient.Get(users, common.ETCD_PREFIX_PLAYERS_INFO, clientv3.WithPrefix())

	t := *user
	log.Println(t)

	if nil != err {
		if ur.myEtcdClient.GetErrorType(err) != error.EcodeKeyNotFound {
			log.Println("Error StatusInternalServerError : ", err)
			response.WriteError(http.StatusInternalServerError, err)
			return
		}
	}

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
