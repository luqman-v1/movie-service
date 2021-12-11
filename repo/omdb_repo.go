package repo

import (
	"context"
	"fmt"
	"log"
	"movie-service/builder"
	"movie-service/entity"
	"movie-service/repo/curl"
	"movie-service/repo/mysql"
	"movie-service/repo/redis"
	pb "movie-service/transport/grpc/proto/movie"
	"movie-service/util"
	"os"

	"golang.org/x/sync/errgroup"
)

var header = map[string]string{
	"Accept": "application/json",
}

type OmDbRepo struct {
	curl    curl.ICurl
	builder builder.OmDb
	movieDB mysql.IMovie
	redis   redis.IRedisDb
}

func NewOmDbRepo() *OmDbRepo {
	return &OmDbRepo{
		curl:    curl.New(os.Getenv("omdb_url"), header),
		builder: builder.NewOmDb(),
		movieDB: mysql.NewMovie(),
		redis:   redis.NewRedis(),
	}
}

type IOmDb interface {
	List(ctx context.Context, request *pb.ListRequest) (*pb.ListResponse, error)
	Detail(ctx context.Context, request *pb.DetailRequest) (*pb.DetailMovie, error)
}

//List get list of movie from IMDB
func (o *OmDbRepo) List(ctx context.Context, request *pb.ListRequest) (resp *pb.ListResponse, err error) {
	bodyBytes, err := o.get("", map[string]string{
		"s":    request.Search,
		"page": util.ItoS(int(request.Page)),
	})
	if err != nil {
		log.Println("error Get", err)
		return
	}
	var list entity.List
	err = list.Decode(bodyBytes)
	if err != nil {
		log.Println("error Decode", err)
		return nil, err
	}
	g, ctx := errgroup.WithContext(ctx)
	err = o.InsertLog(ctx, g, list.Search)
	if err != nil {
		log.Println("error InsertLog with go routine", err)
	}

	resp = &pb.ListResponse{
		Data:       o.builder.BuildResponse(g, list),
		Pagination: o.builder.BuildPagination(request.Page, util.StoI(list.TotalResults)),
	}
	if err = g.Wait(); err != nil {
		log.Println("error on process go routine", err)
	}
	return resp, nil
}

//Detail get detail from IMDB
func (o *OmDbRepo) Detail(ctx context.Context, request *pb.DetailRequest) (resp *pb.DetailMovie, err error) {
	var detail entity.Detail
	key := fmt.Sprintf(redis.GetKey, request.ImdbId)
	get, err := o.redis.Get(ctx, key)
	if err == nil {
		if err := detail.Decode(get); err != nil {
			return nil, util.NotFoundError("IMDb ID")
		}
	} else {
		bodyBytes, err := o.get("", map[string]string{
			"i": request.GetImdbId(),
		})
		if err != nil {
			log.Println("error at Get Detail", err)
			return nil, util.NotFoundError("IMDb ID")
		}
		err = detail.Decode(bodyBytes)
		if err != nil || detail.Response == "False" {
			return nil, util.NotFoundError("IMDb ID")
		}
		if err := o.redis.Set(ctx, key, bodyBytes, redis.DefaultDuration); err != nil {
			log.Println("error at set redis", err)
		}
	}

	return o.builder.BuildDetailResponse(detail), nil
}

func (o *OmDbRepo) get(endpoint string, qParams map[string]string) (b []byte, err error) {
	params := map[string]string{
		"apikey": os.Getenv("omdb_apikey"),
	}
	for k, v := range qParams {
		params[k] = v
	}
	get, err := o.curl.Get(endpoint, params)
	if err != nil {
		return
	}
	return get, nil
}

//InsertLog create log movie with go routine
func (o *OmDbRepo) InsertLog(ctx context.Context, g *errgroup.Group, list []entity.Search) (err error) {
	for _, v := range list {
		value := v
		g.Go(func() error {
			return o.movieDB.Upsert(ctx, value)
		})
	}
	return
}
