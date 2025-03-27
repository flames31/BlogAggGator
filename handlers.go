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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage : %s <username>", cmd.Name)
	}

	usrName := cmd.Args[0]

	dbUser, err := s.db.GetUser(context.Background(), usrName)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	err = s.cfg.SetUser(dbUser.Name)
	if err != nil {
		return fmt.Errorf("could not set user: %w", err)
	}

	fmt.Printf("Switched to user : %v\n", usrName)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage : %s <username>", cmd.Name)
	}

	usrName := cmd.Args[0]

	dbUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: usrName})
	if err != nil {
		return fmt.Errorf("error while creating user: %w", err)
	}
	fmt.Printf("User registered!\n")
	s.cfg.SetUser(dbUser.Name)

	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("All rows removed from the users table")
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		fmt.Printf("* %v", user.Name)
		if user.Name == s.cfg.CurrentUserName {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	rssFeed, err := feed.GetFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(*rssFeed)

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
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

	fmt.Println(feed)
	return nil
}
