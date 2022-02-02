package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"

	"github.com/dandezille/feedbin-export/feedbin"
	"github.com/dandezille/feedbin-export/pinboard"
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
	pinboard := createPinboardClient()

	entries, err := feedbin.GetStarredEntries()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(entries)
	if len(entries) == 0 {
		log.Println("No starred entries")
		return
	}

	for _, entry := range entries {
		err = pinboard.CreateEntry(entry.Url, entry.Title)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = feedbin.Unstar(entries)
	if err != nil {
		log.Fatal(err)
	}
}

func createFeedbinClient() feedbin.Client {
	feedbinUser := utils.ReadEnv("FEEDBIN_USER")
	feedbinPassword := utils.ReadEnv("FEEDBIN_PASSWORD")
	return feedbin.Connect(feedbinUser, feedbinPassword)
}

func createPinboardClient() *pinboard.Client {
	key := utils.ReadEnv("PINBOARD_API_KEY")
	client, err := pinboard.Connect(key)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
