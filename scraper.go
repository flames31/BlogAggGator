package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"html"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/flames31/BlogAggGator/internal/database"
	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(feedUrl string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", feedUrl, nil)
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

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("Couldn't fetch feed")
		return
	}
	_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}
	for _, item := range feedData.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			FeedID:    feed.ID,
			Title:     item.Title,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			Url:         item.Link,
			PublishedAt: publishedAt,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}
