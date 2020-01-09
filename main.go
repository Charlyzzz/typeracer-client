package main

import (
	"context"
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	proto "github.com/Charlyzzz/typeracer-client/typeracer"
	"github.com/eiannone/keyboard"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:8080"
)

func listenKeyStrokes() (<-chan keyboard.Key, error) {
	strokes := make(chan keyboard.Key)
	err := keyboard.Open()
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(strokes)
		for {
			_, key, err := keyboard.GetKey()
			if err != nil {
				log.Fatalf("keyboard get key failed %v", err)
			}
			strokes <- key
		}
	}()
	return strokes, nil
}

func pushMetrics(name string, t *timestamp.Timestamp, stream proto.TypeRacer_SendPlayerMetricsClient) chan<- int {
	metrics := make(chan int)
	go func() {
		defer close(metrics)
		for {
			m := <-metrics
			metric := &proto.PlayerMetrics{Username: name, ConnectionTime: t, StrokesPerMinute: int32(m)}
			err := stream.Send(metric)
			if err != nil {
				log.Fatalf("send player metrics failed %v", err)
			}
		}
	}()
	return metrics
}

func main() {
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("name is required!")
	}
	name := args[0]
	connTime, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		log.Fatalf("timestamp is invalid %v", err)
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	c := proto.NewTypeRacerClient(conn)
	ctx := context.Background()
	stream, err := c.SendPlayerMetrics(ctx)
	if err != nil {
		log.Fatalf("could not send player metrics %v", err)
	}
	metrics := pushMetrics(name, connTime, stream)
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				log.Fatalf("server closed stream")
			}
			if err != nil {
				log.Fatalf("scoreboard receive failed %v", err)
			}
			tm.Clear()
			tm.MoveCursor(0, 0)
			for i, p := range in.Reply {
				player := fmt.Sprintf("#%d => %s: %d\n", i+1, p.Username, p.StrokesPerMinute)
				tm.Print(player)
			}
			tm.Flush()
		}
	}()
	strokes, err := listenKeyStrokes()
	if err != nil {
		log.Fatalf("keyboard open error %v", err)
	}
	var strokeCount int
	ticker := time.NewTicker(500 * time.Millisecond)

	cleanUp := func() {
		stream.CloseSend()
		conn.Close()
		os.Exit(0)
	}

	for {
		select {
		case key := <-strokes:
			if key == keyboard.KeyCtrlC {
				cleanUp()
			}
			strokeCount += 2 * 60
		case <-ticker.C:
			metrics <- strokeCount
			strokeCount = 0
		case <-exit:
			cleanUp()
		}
	}
}
