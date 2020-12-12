package main

import (
  "github.com/joho/godotenv"
  "log"
  "time"

  "github.com/dandezille/feedbin-to-todoist/utils"
  "github.com/dandezille/feedbin-to-todoist/feedbin"
  "github.com/dandezille/feedbin-to-todoist/todoist"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Println("Failed to load .env")
  }

  interval, err := time.ParseDuration(utils.ReadEnv("TICKER_INTERVAL"))
  if err != nil {
    log.Fatal(err.Error())
  }

  ticker := time.NewTicker(interval)
  defer ticker.Stop()

  for ; true; <- ticker.C {
    fetchFeeds()
  }
}

func fetchFeeds() {
  feedbin := feedbin.Connect()
  entries := feedbin.GetStarredEntries()
  log.Println(entries)

  if len(entries) == 0 {
    log.Println("No starred entries")
    return
  }

  todoist := todoist.Connect()
  for _, entry := range entries {
    todoist.CreateEntry(entry.Url)
  }

  feedbin.Unstar(entries)
}
