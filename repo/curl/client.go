package curl

import (
	"github.com/go-resty/resty/v2"
)

type Curl struct {
	Host   string
	Header map[string]string
}

type ICurl interface {
	Get(endpoint string, qParam map[string]string) ([]byte, error)
}

func New(
	host string,
	header map[string]string,
) ICurl {
	return &Curl{
		Host:   host,
		Header: header,
	}
}

func (c *Curl) client() *resty.Request {
	return resty.New().R().
		SetHeaders(c.Header)
}

// Get method get curl
func (c *Curl) Get(endpoint string, qParam map[string]string) ([]byte, error) {
	response, err := c.client().
		SetQueryParams(qParam).
		Get(c.Host + endpoint)
	return response.Body(), err
}
