package repo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"movie-service/entity"
	"movie-service/repo/mocks"
	"movie-service/repo/redis"
	pb "movie-service/transport/grpc/proto/movie"
	"testing"

	"github.com/bxcodec/faker/v3"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/suite"
)

var ctx = context.Background()

type OmDbTestSuite struct {
	suite.Suite
	curl  *mocks.ICurl
	db    *mocks.IMovie
	redis *mocks.IRedisDb
	act   *OmDbRepo
}

func (s *OmDbTestSuite) SetupTest() {
	s.curl = new(mocks.ICurl)
	s.db = new(mocks.IMovie)
	s.redis = new(mocks.IRedisDb)
	s.act = NewOmDbRepo()
	s.act.curl = s.curl
	s.act.movieDB = s.db
	s.act.redis = s.redis
}

func TestOmDbTestSuite(t *testing.T) {
	suite.Run(t, new(OmDbTestSuite))
}

func (s *OmDbTestSuite) AfterTest(_, _ string) {
	s.curl.AssertExpectations(s.T())
	s.db.AssertExpectations(s.T())
	s.redis.AssertExpectations(s.T())
}

func (s *OmDbTestSuite) TestOmDbRepo_List() {
	req := &pb.ListRequest{
		Page:   1,
		Search: "Batman",
	}
	search := []entity.Search{
		{
			Title:  "Batman Year One: A Conversation with DC Comics",
			Year:   "2011",
			ImdbID: "tt2098631",
			Type:   "movie",
			Poster: "https://m.media-amazon.com/images/M/MV5BNTM0M2JlMjUtNTcyNi00NTkzLWI5MmQtNmY1MDJmZDExMDhlXkEyXkFqcGdeQXVyNjM2NDIwMzQ@._V1_SX300.jpg",
		},
	}
	list := entity.List{
		Search:       search,
		TotalResults: "10",
		Response:     "True",
	}
	detail := search[0]
	b, err := json.Marshal(list)

	s.curl.On("Get", "", mock.Anything).Return(b, nil)
	s.db.On("Upsert", mock.Anything, detail).Return(nil)
	response, err := s.act.List(ctx, req)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), detail.Title, response.Data[0].Title)
}

func (s *OmDbTestSuite) TestOmDbRepo_Detail_If() {
	req := &pb.DetailRequest{
		ImdbId: "tt0290453",
	}
	var detail = entity.Detail{}
	_ = faker.FakeData(&detail)
	b, err := detail.Encode()
	key := fmt.Sprintf(redis.GetKey, req.ImdbId)
	s.redis.On("Get", ctx, key).Return(b, nil)
	movie, err := s.act.Detail(ctx, req)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), detail.ImdbID, movie.ImdbID)
}

func (s *OmDbTestSuite) TestOmDbRepo_Detail_Else() {
	req := &pb.DetailRequest{
		ImdbId: "tt0290453",
	}
	var detail = entity.Detail{}
	_ = faker.FakeData(&detail)
	b, err := detail.Encode()
	key := fmt.Sprintf(redis.GetKey, req.ImdbId)
	s.redis.On("Get", ctx, key).Return(nil, errors.New("data empty"))
	s.curl.On("Get", "", mock.Anything).Return(b, nil)
	s.redis.On("Set", ctx, key, b, redis.DefaultDuration).Return(nil)
	movie, err := s.act.Detail(ctx, req)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), detail.ImdbID, movie.ImdbID)
}
