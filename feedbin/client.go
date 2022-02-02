package feedbin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dandezille/feedbin-export/utils"
)

type Client struct {
	user     string
	password string
	client   utils.RestClient
}

type FeedEntry struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

func Connect(user string, password string) Client {
	customiseRequest := func(request *http.Request) {
		request.SetBasicAuth(user, password)
	}

	c := Client{
		user:     user,
		password: password,
		client:   utils.NewRestClient("https://api.feedbin.com/v2/", customiseRequest),
	}
	c.ensureAuthenticated()
	return c
}

func (c *Client) GetStarredEntries() []FeedEntry {
	starred := c.getStarredEntries()
	if len(starred) == 0 {
		return nil
	}

	entries := c.getEntries(starred)
	return entries
}

func (c *Client) Unstar(entries []FeedEntry) {
	var ids []string
	for _, entry := range entries {
		ids = append(ids, strconv.Itoa(entry.Id))
	}

	body := `{"starred_entries": [` + strings.Join(ids, ",") + "]}"
	log.Println(body)
	c.client.Delete("starred_entries.json", body)
}

func (c *Client) ensureAuthenticated() {
	c.client.Get("authentication.json")
}

func (c *Client) getStarredEntries() []int {
	response := c.client.Get("starred_entries.json")

	var starred []int
	err := json.Unmarshal(response, &starred)
	if err != nil {
		log.Fatal(err)
	}

	return starred
}

func (c *Client) getEntries(starred []int) []FeedEntry {
	ids := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(starred)), ","), "[]")
	response := c.client.Get("entries.json?ids=" + ids)

	var entries []FeedEntry
	err := json.Unmarshal(response, &entries)
	if err != nil {
		log.Fatal(err)
	}

	return entries
}
