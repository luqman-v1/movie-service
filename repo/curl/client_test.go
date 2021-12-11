package curl

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurl_Get(t *testing.T) {
	n := New("http://www.omdbapi.com/", nil)
	_, err := n.Get("", map[string]string{
		"apikey": os.Getenv("omdb_apikey"),
		"s":      "Batman",
	})
	assert.Nil(t, err)
}
