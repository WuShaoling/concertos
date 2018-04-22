package executor

import (
	"github.com/concertos/module/common"
	"context"
	"github.com/docker/docker/api/types"
	"log"
	"github.com/docker/docker/client"
)

type DockerExecutor struct {
	myEtcdClient *common.MyEtcdClient
	DockerCli    *client.Client
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
