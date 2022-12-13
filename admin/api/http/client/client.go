package client

import (
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	client *http.Client
	base   string
}

func NewClient(client *http.Client, base string) *Client {
	return &Client{
		client: client,
		base:   base,
	}
}

func (c *Client) Get(pathElem []string) (*http.Response, error) {
	reqURL, err := url.JoinPath(c.base, pathElem...)
	if err != nil {
		return nil, err
	}
	return c.client.Get(reqURL)
}

func (c *Client) Put(pathElem []string, contentType string, body io.Reader) (*http.Response, error) {
	reqURL, err := url.JoinPath(c.base, pathElem...)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, reqURL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.client.Do(req)
}

func (c *Client) Delete(pathElem []string) (*http.Response, error) {
	reqURL, err := url.JoinPath(c.base, pathElem...)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return nil, err
	}
	return c.client.Do(req)
}
