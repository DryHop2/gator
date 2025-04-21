package main

import (
	"fmt"
	"log"

	"github.com/DryHop2/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}

	err = cfg.SetUser("DryHop")
	if err != nil {
		log.Fatalf("could not set user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("could not re-read config: %v", err)
	}

	fmt.Printf("Current config: %+v\n", cfg)
}
