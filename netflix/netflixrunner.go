package netflix

import (
	"context"
	"fmt"
	"log"

	"example.com/assignment/data"
	"go.jonnrb.io/speedtest/fastdotcom"
	"go.jonnrb.io/speedtest/oututil"
	"go.jonnrb.io/speedtest/units"
	"golang.org/x/sync/errgroup"
)

const (
	ExpectedNumberOfServers = 10
)

func RunMe() data.RunResult {
	var client fastdotcom.Client
	ctx := context.Background()
	manifestPtr, _ := fastdotcom.GetManifest(ctx, ExpectedNumberOfServers)

	download(manifestPtr, &client)
	upload(manifestPtr, &client)

	return data.RunResult{}
}

func download(m *fastdotcom.Manifest, client *fastdotcom.Client) float64 {
	ctx := context.Background()

	stream, finalize := proberPrinter(func(s units.BytesPerSecond) string {
		return formatSpeed("Download speed", s)
	})
	speed, err := m.ProbeDownloadSpeed(ctx, client, stream)
	if err != nil {
		log.Fatalf("Error probing download speed: %v", err)
		return -1
	}

	finalize(speed)

	return float64(speed)
}

func upload(m *fastdotcom.Manifest, client *fastdotcom.Client) float64 {
	ctx := context.Background()

	stream, finalize := proberPrinter(func(s units.BytesPerSecond) string {
		return formatSpeed("Upload speed", s)
	})
	speed, err := m.ProbeUploadSpeed(ctx, client, stream)
	if err != nil {
		log.Fatalf("Error probing upload speed: %v", err)
		return -1
	}

	finalize(speed)

	return float64(speed)
}

func proberPrinter(format func(units.BytesPerSecond) string) (
	stream chan units.BytesPerSecond,
	finalize func(units.BytesPerSecond),
) {
	p := oututil.StartPrinting()
	p.Println(format(units.BytesPerSecond(0)))

	stream = make(chan units.BytesPerSecond)
	var g errgroup.Group
	g.Go(func() error {
		for speed := range stream {
			p.Println(format(speed))
			fmt.Println(format(speed))
		}
		return nil
	})

	finalize = func(s units.BytesPerSecond) {
		g.Wait()
		p.Finalize(format(s))
	}
	return
}

func formatSpeed(prefix string, s units.BytesPerSecond) string {
	return fmt.Sprintf("%s: %v", prefix, s)
}
