package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	err := Env("../.env.example")
	assert.Nil(t, err)
}

func TestItoS(t *testing.T) {
	assert.Equal(t, "10", ItoS(10))
}

func TestStoI(t *testing.T) {
	assert.Equal(t, int32(10), StoI("10"))
}
