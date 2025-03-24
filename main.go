package main

import (
	"fmt"
	"log"

	"github.com/flames31/BlogAggGator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Summ went wrong gang: %v", err)
	}
	fmt.Println(cfg)
	err = cfg.SetUser("flames31")
	if err != nil {
		log.Fatalf("Summ went wrong gang: %v", err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Summ went wrong gang: %v", err)
	}
	fmt.Println(cfg)
}
