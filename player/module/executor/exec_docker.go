package executor

import (
	"github.com/concertos/module/common"
	"github.com/concertos/module/entity"
	"context"
	"github.com/docker/docker/api/types"
	"log"
	"bytes"
)

type DockerExecutor struct {
	myEtcdClient *common.MyEtcdClient
}

func (d *DockerExecutor) Install(con *entity.ContainerInfo) string {

	dc := common.GetDockerClient()

	// check if images exist in localhost
	if images, err := dc.ImageList(context.Background(), types.ImageListOptions{}); nil != err {
		log.Println("Error get image list: ", err)
		return err.Error()
	} else {
		for _, v := range images {
			for _, v2 := range v.RepoTags {
				if v2 == con.BaseImage {
					log.Println("image exist")
					return "success"
				}
			}
		}
	}

	// if not, pull from hub
	reader, err := dc.ImagePull(context.Background(), con.BaseImage, types.ImagePullOptions{})
	if nil != err {
		log.Println(err)
	}

	if nil != reader {
		buf := new(bytes.Buffer)
		buf.ReadFrom(reader)
		s := buf.String()
		log.Println("Result form docker client", s)
		return s
	}
	return "Error read from docker client"

	//
	//for {
	//	if nil != reader {
	//		input, _ := ioutil.ReadAll(reader);
	//		res := string(input)
	//		log.Println("----", input)
	//		log.Println("----")
	//		if len(input) == 0 {
	//			break
	//		}
	//		//res := string(input[0:len(input)-1])
	//		//reader.Close()
	//		////
	//		//log.Println("----", res)
	//	} else {
	//		break
	//	}
	//}
	//
	//return "error"

}

func (d *DockerExecutor) Start(con *entity.ContainerInfo) error {

	return nil
}

func (d *DockerExecutor) Stop(con *entity.ContainerInfo) error {

	return nil
}

func (d *DockerExecutor) Remove(con *entity.ContainerInfo) error {

	return nil
}

func GetDockerExecutor() *DockerExecutor {
	return &DockerExecutor{myEtcdClient: common.GetMyEtcdClient()}
}
