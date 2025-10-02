package server

import (
	"context"
	"errors"
	"fmt"
	"time"

	pb "github.com/Lucascluz/memora-proto/gen"
	"github.com/Lucascluz/memora-server/internal/cache"
)

type Server struct {
	pb.UnimplementedMemoraServiceServer

	cache cache.Cache
	conns map[string]string
}

func NewServer() *Server {
	return &Server{
		cache: *cache.NewCache(),
		conns: make(map[string]string),
	}
}

func (s *Server) Connect(ctx context.Context, req *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {

	// check if client is already connected
	clientKey, exists := s.conns[req.ClientIP]
	if exists {
		return &pb.ConnectionResponse{
			Success:   true,
			ClientKey: clientKey,
		}, errors.New("client already connected")
	}

	// generate key for the the client
	clientKey = genKey(req.ClientIP)

	// entry the client and key to the conn map
	s.conns[req.ClientIP] = clientKey

	// return the new client key
	return &pb.ConnectionResponse{
		Success:   true,
		ClientKey: clientKey,
	}, nil
}

func (s *Server) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {

	// verify the clientKey
	if !s.isValidClientKey(req.ClientKey) {
		return &pb.SetResponse{Success: false, Status: "client key not found"}, errors.New("client not connected")
	}

	// set cache entry
	err := s.cache.Set(req.EntryKey, req.Value)
	if err != nil {
		return nil, err
	}

	return &pb.SetResponse{Success: true, Status: "success"}, nil
}

func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {

	// verify the clientKey
	if !s.isValidClientKey(req.ClientKey) {
		return &pb.GetResponse{Status: "client key not found", Value: nil}, errors.New("client not connected")
	}

	// get cache entry
	value, err := s.cache.Get(req.EntryKey)
	if err != nil {
		return &pb.GetResponse{Status: "not found", Value: nil}, nil
	}

	return &pb.GetResponse{Status: "found", Value: value}, nil
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {

	// verify the clientKey
	if !s.isValidClientKey(req.ClientKey) {
		return &pb.DeleteResponse{Found: false, Status: "client key not found"}, errors.New("client not connected")
	}

	// delete cache entry
	err := s.cache.Delete(req.EntryKey)
	if err != nil {
		return &pb.DeleteResponse{Found: false, Status: "not found"}, nil
	}

	return &pb.DeleteResponse{Found: true, Status: "deleted"}, nil
}

// isValidClientKey checks if the provided client key exists in the connections map
func (s *Server) isValidClientKey(clientKey string) bool {
	for _, key := range s.conns {
		if key == clientKey {
			return true
		}
	}
	return false
}

func genKey(ip string) string {
	return fmt.Sprintf("%s-%d", ip, time.Now().UnixNano())
}
