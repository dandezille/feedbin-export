package utils

import (
  "bytes"
  "io"
  "io/ioutil"
  "log"
  "net/http"
)

type RestClient struct {
  url string
  customiseRequest func(*http.Request)
}

func NewRestClient(url string, customiseRequest func(*http.Request)) RestClient {
  return RestClient { 
    url: url,
    customiseRequest: customiseRequest,
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
  c.customiseRequest(request)

  return request
}

func (c *RestClient) Execute(request *http.Request) []byte {
  client := &http.Client{}
  response, err := client.Do(request)
  if err != nil {
    log.Fatal(err)
  }
  defer response.Body.Close()

  if response.StatusCode != 200 {
    log.Fatal(response.Status)
  }

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    log.Fatal(err)
  }

  log.Println("response: " + string(body))
  return body
}

func bodyFromString(data string) io.Reader {
  if data == "" {
    return nil
  } else {
    return bytes.NewBuffer([]byte(data))
  }
}
