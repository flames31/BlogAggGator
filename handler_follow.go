package main

import (
	"context"
	"fmt"
	"time"

	"github.com/flames31/BlogAggGator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("need url as argument")
	}
	feedUrl := cmd.Args[0]

	return followFeed(s, feedUrl)
}

func followFeed(s *state, feedUrl string) error {
	feed, err := s.db.GetFeedByURL(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	currUser, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedFollowRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Feed %v was followed by %v!\n", feedFollowRow.FeedName, feedFollowRow.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
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
