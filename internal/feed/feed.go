package feed

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
)

func GetFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return &RSSFeed{}, err
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	req.Header.Set("User-Agent", "gator")
	resp, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer resp.Body.Close()

	var rssFeedData RSSFeed

	err = xml.Unmarshal(data, &rssFeedData)
	if err != nil {
		return &RSSFeed{}, err
	}

	rssFeedData.Channel.Title = html.UnescapeString(rssFeedData.Channel.Title)
	rssFeedData.Channel.Description = html.UnescapeString(rssFeedData.Channel.Description)
	for i, item := range rssFeedData.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		rssFeedData.Channel.Item[i] = item
	}

	return &rssFeedData, nil
}
