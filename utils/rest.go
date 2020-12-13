package utils

import (
)

type RestClient struct {
  url string
}

func NewRestClient(url string) RestClient {
  return RestClient { 
    url: url,
  }
}

func (c *RestClient) Url(path string) string {
  return c.url + path
}


