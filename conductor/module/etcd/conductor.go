package etcd

import (
	"github.com/concertos/common"
	"sync"
	"github.com/coreos/etcd/clientv3"
	"time"
	"log"
)

type ConductorApi interface {
	//player etcd rest api
	PutPlayer(info *common.PlayerInfo) error
	GetPlayer(id string) (*common.PlayerInfo, error)
	GetAllPlayer() ([]common.PlayerInfo, error)
	DeletePlayer(id string) error

	//user etcd rest api
	PutUser(user *common.UserInfo) error
	GetAllUser() ([]common.UserInfo, error)
	GetUser(id string) (*common.UserInfo, error)
	DeleteUser(id string) error
}

type Conductor struct {
	client clientv3.Client
}

func NewConductor() *Conductor {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   common.GetEtcdPoints(),
		DialTimeout: 2 * time.Second,
	})

	if err != nil {
		log.Fatal("Error: new etcd client error:", err)
		return nil
	}

	conductor := &Conductor{
		client: *etcdClient,
	}
	return conductor
}

var conductor *Conductor
var once sync.Once

func GetConductor() *Conductor {
	once.Do(func() {
		conductor = NewConductor()
	})
	return conductor
}
