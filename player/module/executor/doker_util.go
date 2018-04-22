package executor

import (
	"bytes"
	"context"
	"github.com/docker/docker/api/types"
	"log"
	"errors"
)

func (de *DockerExecutor) PullImage(images string) error {

	// if exist, return
	if de.checkImages(images) == true {
		return nil
	}

	// if not, pull from hub
	if reader, err := de.DockerCli.ImagePull(context.Background(), images, types.ImagePullOptions{}); nil != err {
		log.Println(err)
		return errors.New("Error read from docker client")
	} else if reader == nil {
		log.Println("Reader is nil")
		return errors.New("Error read from docker client")
	} else {
		buf := new(bytes.Buffer)
		buf.ReadFrom(reader)
		s := buf.String()
		log.Println("Result form docker client pull images", s)
		return nil
	}
	return nil
}
