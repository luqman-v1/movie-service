package mysql

import (
	"context"
	"log"
	"movie-service/entity"
)

type Movie struct {
}

func NewMovie() IMovie {
	return &Movie{}
}

type IMovie interface {
	Upsert(ctx context.Context, req entity.Search) error
}

func (l *Movie) Upsert(ctx context.Context, req entity.Search) error {
	var args []interface{}
	args = append(args, req.Title, req.Year, req.ImdbID, req.Type, req.Poster)
	q := QueryUpSertMovie
	_, err := sqlxDB.Exec(sqlxDB.Rebind(q), args...)
	if err != nil {
		log.Println("error at Upsert Movie", err)
		return err
	}
	return nil
}
