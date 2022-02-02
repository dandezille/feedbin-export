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
	feedbin := createFeedbinClient()
	entries := feedbin.GetStarredEntries()
	log.Println(entries)

	if len(entries) == 0 {
		log.Println("No starred entries")
		return
	}

	todoist := createTodoistClient()
	for _, entry := range entries {
		todoist.CreateEntry(entry.Url)
	}

	feedbin.Unstar(entries)
}

func createFeedbinClient() feedbin.Client {
	feedbinUser := utils.ReadEnv("FEEDBIN_USER")
	feedbinPassword := utils.ReadEnv("FEEDBIN_PASSWORD")
	return feedbin.Connect(feedbinUser, feedbinPassword)
}

func createTodoistClient() todoist.Client {
	todoistKey := utils.ReadEnv("TODOIST_API_KEY")
	return todoist.Connect(todoistKey)
}
