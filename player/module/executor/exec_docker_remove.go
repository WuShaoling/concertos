package executor

import (
	"github.com/concertos/module/dns"
	"log"
	"github.com/concertos/module/common"
	"github.com/docker/docker/api/types"
	"context"
	"github.com/concertos/module/entity"
)

func (de *DockerExecutor) Remove(con *entity.ContainerInfo) string {
	log.Println("delete container :", con.Name)

	//从dns删除信息
	log.Println("delete dns info: ", dns.GetDNSApi().Delete(con.Name))

	//删除文件目录 sent to all player

	//停止并删除容器
	de.Stop(con)
	err := de.DockerCli.ContainerRemove(context.Background(), con.Name, types.ContainerRemoveOptions{
/*		RemoveVolumes: true,
		Force:         true,
		RemoveLinks:   true,*/
	})
	if err != nil {
		log.Println("delete error: ",err)
	}

	//从etcd删除信息
	de.myEtcdClient.Delete(common.ETCD_PREFIX_CONTAINER_INFO + con.Name)

	return "delete container " + con.Name + " ok"
}
