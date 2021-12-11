package entity

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bxcodec/faker/v3"
)

func TestList_Decode(t *testing.T) {
	var listData List
	err := faker.FakeData(&listData)
	if err != nil {
		return
	}
	b, _ := json.Marshal(listData)
	var l List
	err = l.Decode(b)
	assert.Nil(t, err)
}

func TestDetail_Decode(t *testing.T) {
	var detailData Detail
	err := faker.FakeData(&detailData)
	if err != nil {
		return
	}
	var d Detail
	b, _ := json.Marshal(detailData)
	err = d.Decode(b)
	assert.Nil(t, err)
}

func TestDetail_Encode(t *testing.T) {
	var detailData Detail
	err := faker.FakeData(&detailData)
	if err != nil {
		return
	}
	_, err = detailData.Encode()
	assert.Nil(t, err)
}
