package openapi

import (
	"log/slog"
	"net/url"
)

type Option func(*Client) error

func WithLogger(l *slog.Logger) Option {
	return func(c *Client) error {
		if l == nil {
			return ErrSlogNil
		}

		c.l = l
		return nil
	}
}

func WithHeader(header, value string) Option {
	return func(c *Client) error {
		c.headers.Add(header, value)
		return nil
	}
}

func WithApiKey(apiKey string) Option {
	return func(c *Client) error {
		return WithHeader("Authorization", "Bearer "+apiKey)(c)
	}
}

func WithProject(project string) Option {
	return func(c *Client) error {
		return WithHeader("OpenAI-Project", project)(c)
	}
}

func WithOrganisation(organisation string) Option {
	return func(c *Client) error {
		return WithHeader("OpenAI-Organizationt", organisation)(c)
	}
}

func WithBaseUrl(u *url.URL) Option {
	return func(c *Client) error {
		c.baseUrl = u
		return nil
	}
}

func WithBaseUrlString(rawUrl string) Option {
	return func(c *Client) error {
		u, err := url.Parse(rawUrl)
		if err != nil {
			return err
		}

		return WithBaseUrl(u)(c)
	}
}
