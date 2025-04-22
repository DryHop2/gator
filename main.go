package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/DryHop2/gator/internal/commands"
	"github.com/DryHop2/gator/internal/config"
	"github.com/DryHop2/gator/internal/database"
	"github.com/DryHop2/gator/internal/state"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	dbQueries := database.New(db)

	appState := &state.State{
		DB:  dbQueries,
		Cfg: cfg,
	}

	if len(os.Args) < 2 {
		fmt.Println("You must provide a command.")
		os.Exit(1)
	}

	cmd := commands.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	cmds := commands.New()
	cmds.Register("login", commands.HandlerLogin)
	cmds.Register("register", commands.HandlerRegister)
	cmds.Register("reset", commands.HandlerReset)
	cmds.Register("users", commands.HandlerUsers)
	cmds.Register("agg", commands.HandlerAgg)

	err = cmds.Run(appState, cmd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
