package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"

	"github.com/dandezille/feedbin-export/feedbin"
	"github.com/dandezille/feedbin-export/todoist"
	"github.com/dandezille/feedbin-export/utils"
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

	for ; true; <-ticker.C {
		fetchFeeds()
	}
}

func fetchFeeds() {
	feedbinUser := utils.ReadEnv("FEEDBIN_USER")
	feedbinPassword := utils.ReadEnv("FEEDBIN_PASSWORD")
	todoistKey := utils.ReadEnv("TODOIST_API_KEY")

	feedbin := feedbin.Connect(feedbinUser, feedbinPassword)
	entries := feedbin.GetStarredEntries()
	log.Println(entries)

	if len(entries) == 0 {
		log.Println("No starred entries")
		return
	}

	todoist := todoist.Connect(todoistKey)
	for _, entry := range entries {
		todoist.CreateEntry(entry.Url)
	}

	feedbin.Unstar(entries)
}
