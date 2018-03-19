package dccp

import (
	"github.com/concertos/module/common"
	"sync"
)

func (d *Dccp) GetIp() string {
	return "192.168.1.1"
}

var once sync.Once
var dccp *Dccp

type Dccp struct {
	MyEctdClient *common.MyEtcdClient
}

func GetDccp() *Dccp {
	once.Do(func() {
		dccp = &Dccp{
			MyEctdClient: common.GetMyEtcdClient(),
		}
	})
	return dccp
}
