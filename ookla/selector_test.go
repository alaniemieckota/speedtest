package ookla

import (
	"testing"

	"go.jonnrb.io/speedtest/speedtestdotnet"
)

func Test_getTopN_WhenMoreServersThanN_ShouldTakeN(t *testing.T) {
	servers := []speedtestdotnet.Server{
		{}, {}, {}, {}, {},
	}
	n := len(servers) / 2
	expectedNumberOfServers := n

	result := getTopN(servers, n)

	if len(result) != n {
		t.Errorf("incorrect result, expected %v, got %v", expectedNumberOfServers, len(result))
	}
}

func Test_getTopN_WhenLessServersThanN_ShouldTakeAllServers(t *testing.T) {
	servers := []speedtestdotnet.Server{
		{}, {}, {},
	}

	expectedNumberOfServers := len(servers)

	n := len(servers) + 2

	result := getTopN(servers, n)

	if len(result) != expectedNumberOfServers {
		t.Errorf("incorrect result, expected %v, got %v", expectedNumberOfServers, len(result))
	}
}

func Test_selectBestServer(t *testing.T) {
	// somehow here a mock for clinet needs to be introduced
	//result := selectBestServer(client)

	// mock for client.Config() to return config
	// mock for client.LoadAllServers

	// expectedBestServer := speedtestdotnet.Server{
	// 	ID:   1,
	// 	Name: "Best server",
	// }

	// if result.ID != expectedBestServer.ID {
	// 	t.Error("given server is not the one I was expecting")
	// }
}
