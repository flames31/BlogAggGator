package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/flames31/BlogAggGator/internal/database"
	"github.com/flames31/BlogAggGator/internal/feed"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	rssFeed, err := feed.GetFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(*rssFeed)

	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return errors.New("we need two args. usage: addfeed <feed_name> <feed_url>")
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: feedName, Url: feedUrl, Name_2: s.cfg.CurrentUserName})
	if err != nil {
		return nil
	}
	fmt.Println("Feed created succesfully!")
	err = followFeed(s, feed, user)
	if err != nil {
		return err
	}
	return nil
}

func handlerGetFeeds(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Printing all feeds in the DB...")
	for _, feed := range feeds {
		user, err := s.db.GetUserByFeedID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Println("------------------------------------")
		fmt.Println("Name : " + feed.Name)
		fmt.Println("URL : " + feed.Url)
		fmt.Println("Created by : " + user.Name)
	}
	return nil
}
