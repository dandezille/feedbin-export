package feedbin

import (
  "log"
  "net/http"
  "io/ioutil"

  "github.com/dandezille/feedbin-to-todoist/utils"
)

func feedbinUrl(path string) string {
  baseUrl := "https://api.feedbin.com/v2/"
  return baseUrl + path
}

func Request(path string) []byte {
  url := feedbinUrl(path)
  log.Println("request: " + url)

  request, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Fatal(err)
  }

  user := utils.ReadEnv("FEEDBIN_USER")
  password := utils.ReadEnv("FEEDBIN_PASSWORD")

  request.SetBasicAuth(user, password)
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

