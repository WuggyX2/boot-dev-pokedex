package main

import (
	"fmt"
	"os"
)

type commandRegistery struct {
	cmds map[string]cliCommand
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func registerCmds() map[string]cliCommand {

	cmds := map[string]cliCommand{
		"exit": {name: "exit", description: "Exit the pokedex", callback: commandExit},
	}

	cmds["help"] = cliCommand{name: "help", description: "Displayes the help message", callback: func() error { return commandHelp(cmds) }}

	return cmds
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commands map[string]cliCommand) error {
  fmt.Println("Welcome to the Pokedex!")
  fmt.Println("Usage:")
  fmt.Println("")

	for _, cmd := range commands {
    fmt.Printf("%s: %s\n", cmd.name, cmd.description)

	}
	return nil
}
