package action

import (
	"context"
	"log"
	"movie-service/repo"
	pb "movie-service/transport/grpc/proto/movie"
)

type List struct {
	repo repo.IOmDb
}

func NewList() *List {
	return &List{
		repo: repo.NewOmDbRepo(),
	}
}

func (l *List) Handler(ctx context.Context, request *pb.ListRequest) (*pb.ListResponse, error) {
	log.Println("request Movie List", request)
	return l.repo.List(ctx, request)
}
