package service

import (
	"log"
)

type Client struct {
	baseUrl string
}

func NewClient(url string) *Client {
	return &Client{url}
}

func (c *Client) GetData(key string) string {
	log.Printf("Getting data from remote service at %s/%s", c.baseUrl, key)
	return key + "=value"
}
