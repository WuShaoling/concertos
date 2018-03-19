package main

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"io/ioutil"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	//res, err := cli.ImageList(ctx, types.ImageListOptions{})
	//
	//for i := 0; i < 1; i++ {
	//	println("Containers --> ", res[i].Containers)
	//	println("Created --> ", res[i].Created)
	//	println("ID --> ", res[i].ID)
	//	for i, j := range res[i].Labels {
	//		println("Labels --> ", i, j)
	//	}
	//	for j := range res[i].RepoTags {
	//		println("RepoTags --> ", res[i].RepoTags[j])
	//	}
	//	for j := range res[i].RepoDigests {
	//		println("RepoDigests --> ", res[i].RepoDigests[j])
	//	}
	//
	//}
	//
	reader, err := cli.ImagePull(ctx, "docker.io/tomcat", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	input, _ := ioutil.ReadAll(reader);

	println(string(input[0:len(input)-1]))
}
