package openapi

import (
	"errors"
	"log/slog"
	"net/http"
	"net/url"
)

var (
	ErrSlogNil = errors.New("provided slog logger was nil")
)

type Client struct {
	l *slog.Logger
	ts http.RoundTripper

	baseUrl *url.URL

	headers http.Header
}

func New(opts ...Option) (*Client, error) {
	c := Client{
		l: slog.Default(),
		ts: http.DefaultTransport,
	}

	for _, f := range opts {
		if err := f(&c); err != nil {
			return nil, err
		}
	}

	return &c, nil
}

func Must(c *Client, err error) (*Client){
	if err!= nil {
		panic(err)
	}

	return c
}
