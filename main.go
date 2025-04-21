package main

import (
	"fmt"
	"os"

	"github.com/DryHop2/gator/internal/commands"
	"github.com/DryHop2/gator/internal/config"
	"github.com/DryHop2/gator/internal/state"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}

	appState := &state.State{Cfg: cfg}

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

	err = cmds.Run(appState, cmd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
