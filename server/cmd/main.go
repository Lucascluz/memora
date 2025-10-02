package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

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
	memoraServer := server.NewServer()
	pb.RegisterMemoraServiceServer(grpcServer, memoraServer)

	// Graceful shutdown support
	go func() {
		log.Println("Server listening on :1212")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	grpcServer.GracefulStop()
	log.Println("Server stopped gracefully")
}
