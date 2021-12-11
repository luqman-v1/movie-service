package builder

import (
	"encoding/json"
	"movie-service/entity"
	pb "movie-service/transport/grpc/proto/movie"

	"golang.org/x/sync/errgroup"
)

type OmDb struct {
}

func NewOmDb() OmDb {
	return OmDb{}
}

var pageSize int32 = 10

func (o *OmDb) BuildPagination(currentPage, totalResult int32) *pb.Pagination {
	return &pb.Pagination{
		PageSize:    pageSize,
		CurrentPage: currentPage,
		TotalPage:   totalResult / pageSize,
		TotalResult: totalResult,
	}
}

func (o *OmDb) BuildResponse(g *errgroup.Group, source entity.List) []*pb.Data {
	results := make([]*pb.Data, len(source.Search))
	for i, v := range source.Search {
		key, value := i, v
		g.Go(func() error {
			results[key] = &pb.Data{
				Title:  value.Title,
				Year:   value.Year,
				ImdbID: value.ImdbID,
				Type:   value.Type,
				Poster: value.Poster,
			}
			return nil
		})
	}
	return results
}

func (o *OmDb) BuildDetailResponse(source entity.Detail) *pb.DetailMovie {
	var resp *pb.DetailMovie
	b, err := source.Encode()
	if err != nil {
		return nil
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil
	}
	return resp
}
