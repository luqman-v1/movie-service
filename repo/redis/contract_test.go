package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"movie-service/entity"
	"movie-service/util"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func init() {
	util.Env("../../.env")
	Conn()
}

var ctx = context.Background()

func getKey() string {
	return fmt.Sprintf(GetKey, "123")
}

func TestRedisDB_Get(t *testing.T) {
	n := NewRedis()
	_, err := n.Get(ctx, getKey())
	log.Println(getKey())
	assert.NotNil(t, err)
}

func TestRedisDB_Set(t *testing.T) {
	n := NewRedis()
	var detail entity.Detail
	faker.FakeData(&detail)
	b, _ := json.Marshal(detail)
	err := n.Set(ctx, getKey(), b, DefaultDuration)
	assert.Nil(t, err)

	_, err = n.Get(ctx, getKey())
	assert.Nil(t, err)

	err = n.Del(ctx, getKey())
	assert.Nil(t, err)
}

func TestRedisDB_Del(t *testing.T) {
	n := NewRedis()
	err := n.Del(ctx, getKey())
	assert.Nil(t, err)
}
