package grpc

import (
	"context"
	"movie-service/action"
	pb "movie-service/transport/grpc/proto/movie"
)

type Server struct{}

func (s *Server) List(ctx context.Context, request *pb.ListRequest) (*pb.ListResponse, error) {
	return action.NewList().Handler(ctx, request)
}

func (s *Server) Detail(ctx context.Context, request *pb.DetailRequest) (*pb.DetailMovie, error) {
	return action.NewDetail().Handler(ctx, request)
}

func NewServer() *Server {
	return &Server{}
}
