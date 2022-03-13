package ookla

import (
	"context"
	"fmt"
	"log"

	"example.com/assignment/data"
	"go.jonnrb.io/speedtest/oututil"
	"go.jonnrb.io/speedtest/speedtestdotnet"
	"go.jonnrb.io/speedtest/units"
	"golang.org/x/sync/errgroup"
)

func RunMe() data.RunResult {
	var client speedtestdotnet.Client

	theServer := selectBestServer(&client)
	downloadSpeedBytes := download(&client, theServer)
	uploadSpeedBytes := upload(&client, theServer)

	return data.RunResult{
		DownloadMb: downloadSpeedBytes / float64(units.MBps),
		UploadMb:   uploadSpeedBytes / float64(units.Mbps),
	}
}

func download(client *speedtestdotnet.Client, server speedtestdotnet.Server) float64 {
	ctx := context.Background()

	stream, finalize := proberPrinter(func(s units.BytesPerSecond) string {
		return formatSpeed("Upload speed", s)
	})

	speed, err := server.ProbeDownloadSpeed(ctx, client, stream)
	if err != nil {
		log.Fatalf("Error probing download speed: %v", err)
		return 0
	}

	finalize(speed)

	return float64(speed)
}

func upload(client *speedtestdotnet.Client, server speedtestdotnet.Server) float64 {
	ctx := context.Background()

	stream, finalize := proberPrinter(func(s units.BytesPerSecond) string {
		return formatSpeed("Upload speed", s)
	})
	speed, err := server.ProbeUploadSpeed(ctx, client, stream)
	if err != nil {
		log.Fatalf("Error probing upload speed: %v", err)
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
