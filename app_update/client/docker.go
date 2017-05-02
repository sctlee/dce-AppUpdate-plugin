package client

import (
    "github.com/docker/cli/client"
)

var dockerCli *client.Client

// GetDockerClient ...
func GetDockerClient() *client.Client {
    dockerCli, err := client.NewEnvClient()
    if err != nil {
        panic(err)
    }
    return dockerCli
}
