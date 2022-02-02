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
	Id    int    `json:"id"`
	Url   string `json:"url"`
	Title string `json:"title"`
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

func (c *Client) GetStarredEntries() ([]FeedEntry, error) {
	starred, err := c.getStarredEntries()
	if err != nil {
		return nil, err
	}

	if len(starred) == 0 {
		return nil, nil
	}

	return c.getEntries(starred)
}

func (c *Client) Unstar(entries []FeedEntry) error {
	var ids []string
	for _, entry := range entries {
		ids = append(ids, strconv.Itoa(entry.Id))
	}

	body := `{"starred_entries": [` + strings.Join(ids, ",") + "]}"
	log.Println(body)

	_, err := c.client.Delete("starred_entries.json", body)
	return err
}

func (c *Client) ensureAuthenticated() {
	_, err := c.client.Get("authentication.json")
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Client) getStarredEntries() ([]int, error) {
	response, err := c.client.Get("starred_entries.json")
	if err != nil {
		return nil, err
	}

	var starred []int
	err = json.Unmarshal(response, &starred)
	if err != nil {
		return nil, err
	}

	return starred, nil
}

func (c *Client) getEntries(starred []int) ([]FeedEntry, error) {
	ids := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(starred)), ","), "[]")
	response, err := c.client.Get("entries.json?ids=" + ids)
	if err != nil {
		return nil, err
	}

	var entries []FeedEntry
	err = json.Unmarshal(response, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
