package todoist

import (
  "log"
  "io"
  "bytes"
  "net/http"
  "io/ioutil"

  "github.com/dandezille/feedbin-to-todoist/utils"
)

type TodoistClient struct {
}

func Connect() TodoistClient {
  ensureAuthenticated()
  return TodoistClient{}
}

func (c *TodoistClient) CreateEntry(content string) {
  post("tasks", `{ "content": "` + content + `" }`)
}

func ensureAuthenticated() {
  get("projects")
}

func url(path string) string {
  baseUrl := "https://api.todoist.com/rest/v1/"
  return baseUrl + path
}

func get(path string) {
  request("GET", path, "")
}

func post(path string, data string) {
  request("POST", path, data)
}

func request(method string, path string, data string) []byte {
  url := url(path)
  log.Println("request: " + url)

  request, err := http.NewRequest(method, url, getBody(data))
  if err != nil {
    log.Fatal(err)
  }

  key := utils.ReadEnv("TODOIST_API_KEY")
  request.Header.Add("Authorization", "Bearer " + key)
  request.Header.Add("Content-Type", "application/json")

  client := &http.Client{}
  response, err := client.Do(request)
  if err != nil {
    log.Fatal(err)
  }
  defer response.Body.Close()

  if response.StatusCode != 200 {
    log.Fatal("request for " + url + " status " + response.Status)
  }

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    log.Fatal(err)
  }

  log.Println("response: " + string(body))
  return body
}

func getBody(data string) io.Reader {
  if data == "" {
    return nil
  } else {
    return bytes.NewBuffer([]byte(data))
  }
}
