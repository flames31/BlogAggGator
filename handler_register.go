package main

import (
	"context"
	"fmt"
	"time"

	"github.com/flames31/BlogAggGator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage : %s <username>", cmd.Name)
	}

	usrName := cmd.Args[0]

	dbUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: usrName})
	if err != nil {
		return fmt.Errorf("error while creating user: %w", err)
	}
	s.cfg.SetUser(dbUser.Name)

	return nil
}
