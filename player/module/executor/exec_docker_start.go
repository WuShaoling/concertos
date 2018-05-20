package executor

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/concertos/module/common"
	"log"
	"strings"
	"time"
	"github.com/concertos/module/entity"
	"github.com/concertos/module/nfs"
	"github.com/concertos/module/util"
	"github.com/concertos/player/module/manager"
	"github.com/concertos/module/dns"
)

func (de *DockerExecutor) Start(con *entity.ContainerInfo) (string) {

	log.Println("start container : ", con.Name)

	// 如果启动过
	if cons, err := de.DockerCli.ContainerList(context.Background(), types.ContainerListOptions{All: true}); err != nil {
		log.Println(err)
		return err.Error()
	} else {
		for _, v := range cons {
			name := strings.Split(v.Names[0], "/")[1]
			log.Println(name)
			if name == con.Name {
				timeout, _ := time.ParseDuration("30s")
				if err := de.DockerCli.ContainerRestart(context.Background(), v.ID, &timeout); nil != err {
					return err.Error()
				}
				manager.GetManage().ContainerManager.Register(con.Name)

				if con.Port == 5678{
					util.ExecShell("docker restart "+con.Name+"-pulsar")
				}

				return "start success"
			}
		}
	}

	// 检查是否有镜像
	if err := de.PullImage(con.BaseImage); nil != err {
		return err.Error()
	}
	// 创建应用根目录
	if err := nfs.GetNFSApi().CreateContainerRootPath(con.Name); err != nil {
		return err.Error()
	}

	log.Println(common.NFS_MOUNT_LOCAL_PATH + con.Name + ":" + common.NFS_MOUNT_CONTAINER_PATH)

	// create new container
	resp, err := de.DockerCli.ContainerCreate(context.Background(), &container.Config{
		Image: con.BaseImage,
		Volumes: map[string]struct{}{
			common.NFS_MOUNT_CONTAINER_PATH: struct{}{},
		},
	}, &container.HostConfig{
		Binds: []string{common.NFS_MOUNT_LOCAL_PATH + con.Name + ":" + common.NFS_MOUNT_CONTAINER_PATH},
	}, &network.NetworkingConfig{
	}, con.Name)
	if err != nil {
		return err.Error()
	}
	//start
	if err := de.DockerCli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{}); err != nil {
		return err.Error()
	}

	//get container info
	inspect, err := de.DockerCli.ContainerInspect(context.Background(), resp.ID)
	if nil != err {
		return err.Error()
	}
	log.Println(inspect.NetworkSettings.IPAddress)

	// update info and put to etcd
	con.State = common.CONTAINER_STATE_RUNNING
	con.Id = resp.ID
	con.Ip = inspect.NetworkSettings.IPAddress
	de.myEtcdClient.Put(common.ETCD_PREFIX_CONTAINER_INFO+con.Name, string(util.MyJsonMarshal(*con)))

	// register to manage module
	manager.GetManage().ContainerManager.Register(con.Name)

	// update dns module
	dns.GetDNSApi().Add(con.Name, con.Ip)

	log.Println("start container " + con.Name + " ok")

	if con.Port == 5678 {
		util.ExecShell("docker run --name=" + con.Name + "-pulsar -d --net=container:" + con.Name + " daocloud.io/shaoling/concertos-pulsar:master")
	}
	return resp.ID
}
