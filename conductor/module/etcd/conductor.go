package etcd

import (
	"github.com/concertos/common"
	"sync"
	"github.com/coreos/etcd/clientv3"
)

//type ConductorApi interface {
//	//player etcd rest api
//	PutPlayer(info *common.PlayerInfo) error
//	PutPlayerId(id string) error
//	GetPlayer(id string) (*common.PlayerInfo, error)
//	GetAllPlayer() ([]common.PlayerInfo, error)
//	DeletePlayer(id string) error
//
//	//user etcd rest api
//	PutUser(user *common.UserInfo) error
//	GetAllUser() ([]common.UserInfo, error)
//	GetUser(id string) (*common.UserInfo, error)
//	DeleteUser(id string) error
//}

type Conductor struct {
	ClientV3    *clientv3.Client
	MyEtcdClent *common.MyEtcdClient
}

var conductor *Conductor
var once sync.Once

func GetConductor() *Conductor {
	once.Do(func() {
		conductor = &Conductor{
			MyEtcdClent: common.GetMyEtcdClient(),
			ClientV3:    common.GetClientV3(),
		}
	})
	return conductor
}
