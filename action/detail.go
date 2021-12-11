package action

import (
	"context"
	"log"
	"movie-service/repo"
	pb "movie-service/transport/grpc/proto/movie"
)

type Detail struct {
	repo repo.IOmDb
}

func NewDetail() *Detail {
	return &Detail{
		repo: repo.NewOmDbRepo(),
	}
}

func (d *Detail) Handler(ctx context.Context, request *pb.DetailRequest) (*pb.DetailMovie, error) {
	log.Println("request Movie Detail", request)
	return d.repo.Detail(ctx, request)
}
