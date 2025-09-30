package main

import (
	"log"
	"net"

	pb "github.com/Lucascluz/memora-proto/gen"
	"github.com/Lucascluz/memora-server/internal/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":1212")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMemoraServiceServer(grpcServer, server.NewServer())
	log.Println("Server listening on :1212")
	grpcServer.Serve(lis)
}
