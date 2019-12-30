package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"

	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	proto "github.com/Charlyzzz/typeracer-client/typeracer"
	tm "github.com/buger/goterm"
	"google.golang.org/grpc"
)

const (
	address     = "127.0.0.1:8080"
	rate        = 5000
	playerCount = 10
)

type player struct {
	n string
	r int
	t *timestamp.Timestamp
}

func pushMetrics(stream proto.TypeRacer_SendPlayerMetricsClient) chan<- player {
	players := make(chan player)
	go func() {
		defer close(players)
		for {
			p := <-players
			//rand := rand.Intn(10)
			metric := &proto.PlayerMetrics{Username: p.n, StrokesPerMinute: int32(p.r), ConnectionTime: p.t}
			err := stream.Send(metric)
			if err != nil {
				log.Fatalf("send player metrics failed %v", err)
			}
		}
	}()
	return players
}

func main() {
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
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
	metrics := pushMetrics(stream)
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

	const rateDiff = rate / playerCount / 2

	for i := 0; i <= playerCount; i++ {
		go func(n int) {
			ticker := time.NewTicker(500 * time.Millisecond)
			p := player{n: "erwin" + strconv.Itoa(n+1), r: rate - n*rateDiff, t: connTime}
			for {
				<-ticker.C
				metrics <- p
			}
		}(i)
	}
	<-exit
	stream.CloseSend()
	conn.Close()
	os.Exit(0)
}
