package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type RestClient struct {
	url              string
	customiseRequest func(*http.Request)
}

func NewRestClient(url string, customiseRequest func(*http.Request)) RestClient {
	return RestClient{
		url:              url,
		customiseRequest: customiseRequest,
	}
}

func (c *RestClient) Get(path string) ([]byte, error) {
	request, err := c.newRequest("GET", path, "")
	if err != nil {
		return nil, err
	}

	return c.execute(request)
}

func (c *RestClient) Post(path string, data string) ([]byte, error) {
	request, err := c.newRequest("POST", path, data)
	if err != nil {
		return nil, err
	}

	return c.execute(request)
}

func (c *RestClient) Delete(path string, data string) ([]byte, error) {
	request, err := c.newRequest("DELETE", path, data)
	if err != nil {
		return nil, err
	}

	return c.execute(request)
}

func (c *RestClient) newRequest(method string, path string, data string) (*http.Request, error) {
	url := c.url + path
	log.Println(method + ": " + url + " " + data)

	request, err := http.NewRequest(method, url, bodyFromString(data))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	c.customiseRequest(request)

	return request, nil
}

func (c *RestClient) execute(request *http.Request) ([]byte, error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected status code: %v", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	log.Println("response: " + string(body))
	return body, nil
}

func bodyFromString(data string) io.Reader {
	if data == "" {
		return nil
	} else {
		return bytes.NewBuffer([]byte(data))
	}
}
