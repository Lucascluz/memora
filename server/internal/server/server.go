package server

import (
	"context"

	pb "github.com/Lucascluz/memora-proto/gen"
	"github.com/Lucascluz/memora-server/internal/cache"
)

type Server struct {
	pb.UnimplementedMemoraServiceServer

	cache cache.Cache
}

func NewServer() *Server {
	return &Server{
		cache: *cache.NewCache(),
	}
}

func (s *Server) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {

	// set cache entry
	err := s.cache.Set(req.Key, req.Value)
	if err != nil {
		return nil, err
	}

	return &pb.SetResponse{Success: true, Status: "success"}, nil
}

func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {

	// get cache entry
	value, err := s.cache.Get(req.Key)
	if err != nil {
		return &pb.GetResponse{Status: "not found", Value: nil}, nil
	}

	return &pb.GetResponse{Status: "found", Value: value}, nil
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {

	// delete cache entry
	err := s.cache.Delete(req.Key)
	if err != nil {
		return &pb.DeleteResponse{Found: false, Status: "not found"}, nil
	}

	return &pb.DeleteResponse{Found: true, Status: "deleted"}, nil
}
