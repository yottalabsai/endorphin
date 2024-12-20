package main

import (
	"fmt"
	pb "github.com/yottalabsai/endorphin/pkg/services/synapse"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"math/rand"
	"net"
	"time"
)

type ConnectedClient struct {
	ClientId string
	Stream   pb.SynapseService_CallServer
}

type SimpleSynapseServiceServer struct {
}

func (SimpleSynapseServiceServer) Call(stream pb.SynapseService_CallServer) error {
	// 获取元数据
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok || len(md["authorization"]) == 0 {
		return fmt.Errorf("missing authorization")
	}

	go func() {
		for {
			messageId := fmt.Sprintf("%d", rand.Intn(10000))
			err := stream.Send(&pb.YottaLabsStream{MessageId: messageId})
			if err != nil {
				log.Fatalf("server send error: %v", err)
				return
			}
			log.Printf("client send, messageId: %s", messageId)
			time.Sleep(15 * time.Second)
		}
	}()
	for {
		recv, err := stream.Recv()
		if err != nil {
			log.Fatalf("server recv error: %v", err)

		}
		log.Printf("server received, messageId: %s\n", recv.MessageId)
	}
	return nil
}

func main() {

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSynapseServiceServer(grpcServer, &SimpleSynapseServiceServer{})
	log.Println("Server is running on port :50051")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
