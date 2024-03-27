package yelp

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"os"
	"strings"
)

type Client struct {
	id    string
	auth  string
	url   string
	Debug bool
}

func New(config Config) *Client {
	return &Client{
		id:   config.ClientID,
		auth: "Bearer " + config.ApiKey,
		url:  strings.TrimRight(config.Url, "/ "),
	}
}

func (c *Client) fullURL(uri string) string {
	return fmt.Sprintf("%s/%s", c.url, strings.TrimLeft(uri, "/ "))
}

func (c *Client) Get(uri string, model, receiver any) error {
	req, err := http.NewRequest("GET", c.fullURL(uri), nil)
	if err != nil {
		return fmt.Errorf("http.NewRequest; %s; %v", uri, err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.auth)

	q, err := query.Values(model)
	if err != nil {
		return fmt.Errorf("query.Values; %s; %v", uri, err)
	}

	req.URL.RawQuery = q.Encode()

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("client.Do; %s; %v", uri, err)
	}

	return c.parseResponse(res, receiver, uri)
}

func (c *Client) parseResponse(res *http.Response, receiver any, uri string) error {
	out, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("io.ReadAll; %s; %v", uri, err)
	}

	if res.StatusCode == 404 {
		return fmt.Errorf("%s; %s", uri, res.Status)
	}

	if res.StatusCode > 300 {
		return fmt.Errorf("%s; %s; %s", uri, res.Status, string(out))
	}

	if c.Debug {
		_ = os.WriteFile("debug.json", out, 0644)
	}

	err = json.Unmarshal(out, receiver)
	if err != nil {
		return fmt.Errorf("json.Unmarshal; %s; %v", uri, err)
	}

	return nil
}
