package ookla

import (
	"context"

	"go.jonnrb.io/speedtest/speedtestdotnet"
)

const (
	NClosestServers = 3
)

func selectBestServer(client *speedtestdotnet.Client) speedtestdotnet.Server {
	ctx := context.Background()
	config, _ := client.Config(ctx)
	servers, _ := client.LoadAllServers(ctx)
	speedtestdotnet.SortServersByDistance(servers, config.Coordinates)
	closestServers := getTopN(servers, NClosestServers)

	return closestServers[0]
}

func getTopN(servers []speedtestdotnet.Server, n int) []speedtestdotnet.Server {
	if len(servers) > n {
		return servers[:n]
	} else {
		return servers
	}
}
