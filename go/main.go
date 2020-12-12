package main

import (
  "github.com/joho/godotenv"
  "log"

  "github.com/dandezille/feedbin-to-todoist/feedbin"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal(err)
  }

  entries := feedbin.GetStarredEntries()
  log.Println(entries)
}
