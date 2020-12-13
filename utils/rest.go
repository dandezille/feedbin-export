package utils

import (
  "bytes"
  "io"
  "log"
  "net/http"
)

type RestClient struct {
  url string
}

func NewRestClient(url string) RestClient {
  return RestClient { 
    url: url,
  }
}

func (c *RestClient) NewRequest(method string, path string, data string) *http.Request {
  url := c.url + path
  log.Println("GET: " + url)

  request, err := http.NewRequest(method, url, bodyFromString(data))
  if err != nil {
    log.Fatal(err)
  }

  request.Header.Add("Content-Type", "application/json")

  return request
}

func bodyFromString(data string) io.Reader {
  if data == "" {
    return nil
  } else {
    return bytes.NewBuffer([]byte(data))
  }
}
