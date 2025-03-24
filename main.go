package main

import (
	"fmt"
	"log"
	"os"

	"github.com/flames31/BlogAggGator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Summ went wrong gang: %v", err)
	}
	programState := &state{
		cfg: &cfg,
	}

	cmds := commands{
		regsiteredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage : cli <command> [args...]")
		return
	}

	fmt.Println(os.Args)
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
