package main

import (
  "github.com/joho/godotenv"
  "log"

  "github.com/dandezille/feedbin-to-todoist/feedbin"
  "github.com/dandezille/feedbin-to-todoist/todoist"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal(err)
  }

  feedbin := feedbin.Connect()
  entries := feedbin.GetStarredEntries()
  log.Println(entries)

  todoist := todoist.Connect()
  for _, entry := range entries {
    todoist.CreateEntry(entry.Url)
  }
}
