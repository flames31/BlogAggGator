package main

import (
	"context"
	"fmt"
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
