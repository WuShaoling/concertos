package etcd

//type ConductorApi interface {
//	//player common rest api
//	PutPlayer(info *common.PlayerInfo) error
//	PutPlayerId(id string) error
//	GetPlayer(id string) (*common.PlayerInfo, error)
//	GetAllPlayer() ([]common.PlayerInfo, error)
//	DeletePlayer(id string) error
//
//	//user common rest api
//	PutUser(user *common.UserInfo) error
//	GetAllUser() ([]common.UserInfo, error)
//	GetUser(id string) (*common.UserInfo, error)
//	DeleteUser(id string) error
//}
//
//type Conductor struct {
//	MyEtcdClent *common.myEtcdClient
//	RestApi     *restapi.RestApi
//}
//
//var conductor *Conductor
//var once sync.Once
//
//func GetConductor() *Conductor {
//	once.Do(func() {
//		conductor = &Conductor{
//			MyEtcdClent: common.GetMyEtcdClient(),
//			RestApi:     restapi.GetRestApi(),
//		}
//	})
//	return conductor
//}
