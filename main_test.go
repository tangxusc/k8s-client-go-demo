package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"testing"
)

func TestDockerClient(t *testing.T) {
	cli, e := client.NewEnvClient()
	if e != nil {
		println(e.Error())
		return
	}
	version, e := cli.ServerVersion(context.TODO())
	if e != nil {
		println(e.Error())
		return
	}
	fmt.Println(version)
	body, e := cli.ContainerCreate(context.TODO(), &container.Config{
		Image: "tomcat:8",
	}, nil, nil, "test-container")
	if e != nil {
		println(e.Error())
		return
	}
	fmt.Println(body)

	e = cli.ContainerStart(context.TODO(), body.ID, types.ContainerStartOptions{})
	if e != nil {
		println(e.Error())
		return
	}
}
