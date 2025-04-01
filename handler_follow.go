package main

import (
	"context"
	"fmt"
	"time"

	"github.com/flames31/BlogAggGator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("need url as argument")
	}
	feedUrl := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	return followFeed(s, feed, user)
}

func followFeed(s *state, feed database.Feed, user database.User) error {

	feedFollowRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Feed %v was followed by %v!\n", feedFollowRow.FeedName, feedFollowRow.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsByUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	fmt.Printf("Current user %v is following below feeds:\n", s.cfg.CurrentUserName)
	for _, feedFollow := range feedFollows {
		feed, err := s.db.GetFeedByID(context.Background(), feedFollow.FeedID)
		if err != nil {
			return err
		}
		fmt.Printf("%v\n", feed.Name)
	}
	return nil
}
