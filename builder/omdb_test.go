package builder

import (
	"context"
	"movie-service/entity"
	"testing"

	"github.com/bxcodec/faker/v3"

	"golang.org/x/sync/errgroup"

	"github.com/stretchr/testify/assert"
)

var ctx = context.Background()

func TestOmDb_BuildPagination(t *testing.T) {
	var currentPage int32 = 1
	var totalResult int32 = 200
	b := NewOmDb()
	pagination := b.BuildPagination(currentPage, totalResult)
	assert.Equal(t, int32(20), pagination.TotalPage)
	assert.Equal(t, totalResult, pagination.TotalResult)
}

func TestOmDb_BuildResponse(t *testing.T) {
	b := NewOmDb()
	g, _ := errgroup.WithContext(ctx)
	var list entity.List
	err := faker.FakeData(&list)
	if err != nil {
		return
	}
	results := b.BuildResponse(g, list)
	err = g.Wait()
	if err != nil {
		return
	}
	for k, v := range results {
		assert.Equal(t, v.Title, list.Search[k].Title)
	}
}

func TestOmDb_BuildDetailResponse(t *testing.T) {
	b := NewOmDb()
	var detail entity.Detail
	err := faker.FakeData(&detail)
	if err != nil {
		return
	}
	result := b.BuildDetailResponse(detail)
	assert.Equal(t, result.Title, detail.Title)
}
