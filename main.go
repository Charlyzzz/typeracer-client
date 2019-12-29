package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	proto "github.com/Charlyzzz/typeracer-client/grpc"
	"github.com/eiannone/keyboard"
	"google.golang.org/grpc"
)

const (
	address = "10.0.1.5:8080"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Panicf("Name is required!")
	}
	name := args[0]

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewTypeRacerClient(conn)

	err2 := keyboard.Open()
	if err2 != nil {
		panic(err2)
	}
	defer keyboard.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
	defer cancel()
	stream, err := c.SendPlayerMetrics(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	exit := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(exit)
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got message %s", in.Reply[0].Username)
		}
	}()
	strokes := make(chan struct{})
	go func() {
		for {
			_, key, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			} else if key == keyboard.KeyCtrlC {
				close(exit)
			}
			strokes <- struct{}{}
		}
	}()
	metrics := make(chan int)
	go func() {
		for {
			m := <-metrics
			metric := &proto.PlayerMetrics{Username: name, StrokesPerMinute: int32(m)}
			stream.Send(metric)
		}
	}()
	strokeCount := 0
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-strokes:
			strokeCount += 1
		case <-ticker.C:
			metrics <- strokeCount * 2 * 60
			strokeCount = 0
		case <-exit:
			stream.CloseSend()
			os.Exit(0)
		}
	}
}
