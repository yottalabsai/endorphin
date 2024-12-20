package main

import (
	"context"
	"fmt"
	pb "github.com/yottalabsai/endorphin/pkg/services/synapse"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"math/rand"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewSynapseServiceClient(conn)
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", "valid-token"))
	//ctx := context.Background()
	stream, err := client.Call(ctx)
	if err != nil {
		log.Fatalf("error creating stream: %v", err)
		return
	}

	go func() {
		for {
			recv, err := stream.Recv()
			if err != nil {
				log.Panicf("error recv message: %v", err)
				return
			}
			log.Printf("client received, messageId: %s", recv.MessageId)
		}
	}()

	go func() {
		for {
			messageId := fmt.Sprintf("%d", rand.Intn(10000))
			req := &pb.YottaLabsStream{MessageId: messageId}
			if err := stream.Send(req); err != nil {
				log.Panicf("error send message: %v", err)
				return
			}
			log.Printf("client send, messageId: %s", messageId)
			time.Sleep(10 * time.Second)
		}
	}()

	time.Sleep(1200 * time.Second)

	err = stream.CloseSend()
	if err != nil {
		return
	}
}
