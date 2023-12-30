package httptesting

import "github.com/go-resty/resty/v2"

type (
	LibraryClient interface {
		GetBooksEndpoint() ([]byte, error)
	}

	client struct {
		req  *resty.Client
		base string
	}
)

func New(base string) *client {
	return &client{
		req:  resty.New(),
		base: base,
	}
}

func (c *client) GetBooksEndpoint() ([]byte, error) {
	resp, err := c.req.R().EnableTrace().Get(c.base + "/foo/bar")
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}
