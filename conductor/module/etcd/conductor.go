package etcd

import (
	"github.com/concertos/common"
	"sync"
	"github.com/coreos/etcd/client"
	"time"
	"log"
)

type ConductorApi interface {
	//player etcd api
	PlayerExpire(info *common.PlayerInfo) error
	SetPlayer(info *common.PlayerInfo) error
	DeletePlayer(info *common.PlayerInfo) error
	GetPlayer(id string) *common.PlayerInfo

	//user etcd api
	AddUser(info *common.UserInfo) error
	DeleteUser(info *common.UserInfo) error
	UpdateUser(info *common.UserInfo) error
}

type Conductor struct {
	KeysAPI client.KeysAPI
}

func NewConductor() *Conductor {
	cfg := client.Config{
		Endpoints:               common.GetEtcdPoints(),
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := client.New(cfg)
	if err != nil {
		log.Fatal("Error: cannot connect to etcd:", err)
		return nil
	}

	conductor := &Conductor{
		KeysAPI: client.NewKeysAPI(etcdClient),
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
