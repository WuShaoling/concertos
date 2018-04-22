package executor

import (
	"time"
	"context"
	"log"
	"github.com/concertos/module/common"
	"github.com/concertos/module/util"
	"github.com/concertos/module/entity"
	"github.com/concertos/player/module/manager"
)

func (de *DockerExecutor) Stop(con *entity.ContainerInfo) string {
	log.Println("stop container : " + con.Name)
	// stop container
	timeout, _ := time.ParseDuration("60s")
	err := de.DockerCli.ContainerStop(context.Background(), con.Name, &timeout)
	if err != nil {
		log.Println("stop container: ", err)
		return err.Error()
	}

	// update container state stored in etcd
	con.State = common.CONTAINER_STATE_STOPPED
	de.myEtcdClient.Put(common.ETCD_PREFIX_CONTAINER_INFO+con.Name, string(util.MyJsonMarshal(con)))

	manager.GetManage().ContainerManager.Remove(con.Name)

	log.Println("stop container " + con.Name + " ok")

	return "stop container ok"
}
