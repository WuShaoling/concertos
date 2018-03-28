package executor

import (
	"github.com/concertos/module/common"
	"github.com/concertos/module/entity"
	"context"
	"github.com/docker/docker/api/types"
	"log"
	"bytes"
	"github.com/docker/docker/client"
)

type DockerExecutor struct {
	myEtcdClient *common.MyEtcdClient
	DockerCli    *client.Client
}

func (de *DockerExecutor) Install(con *entity.ContainerInfo) string {
	if de.checkImages(con.BaseImage) == true {
		return "success"
	}

	// if not, pull from hub
	if reader, err := de.DockerCli.ImagePull(context.Background(), con.BaseImage, types.ImagePullOptions{}); nil != err {
		log.Println(err)
		return "Error read from docker client"
	} else if reader == nil {
		log.Println("Reader is nil")
		return "Error read from docker client"
	} else {
		buf := new(bytes.Buffer)
		buf.ReadFrom(reader)
		s := buf.String()
		log.Println("Result form docker client", s)
		return s
	}
}

func (de *DockerExecutor) Start(con *entity.ContainerInfo) string {
	//de.DockerCli.Container.ContainerStart(context.Background(), )

	return ""
}

func (de *DockerExecutor) Stop(con *entity.ContainerInfo) string {

	return ""
}

func (de *DockerExecutor) Remove(con *entity.ContainerInfo) string {

	return ""
}

func (de *DockerExecutor) checkImages(id string) bool {
	// check if images exist in localhost
	if images, err := de.DockerCli.ImageList(context.Background(), types.ImageListOptions{}); nil != err {
		log.Println("Error check image list: ", err)
		return false
	} else {
		for _, v := range images {
			for _, v2 := range v.RepoTags {
				if v2 == id {
					log.Println("image exist")
					return true
				}
			}
		}
	}
	return false
}

func GetDockerExecutor() *DockerExecutor {
	return &DockerExecutor{myEtcdClient: common.GetMyEtcdClient(), DockerCli: common.GetDockerClient()}
}
