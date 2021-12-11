package mysql

import (
	"context"
	"movie-service/entity"
	"movie-service/util"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bxcodec/faker/v3"
)

var ctx = context.Background()

func init() {
	util.Env("../../.env")
	Conn()
}

func TestMovie_Upsert(t *testing.T) {
	m := NewMovie()
	var search entity.Search
	faker.FakeData(&search)
	err := m.Upsert(ctx, search)
	assert.Nil(t, err)
}
