package main

import (
	"fmt"
	"os"

	"github.com/WuggyX2/boot-dev-pokedex/pokedex"
)

type handler struct {
	cfg *config
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func registerCmds(cfg *config) map[string]cliCommand {
	handler := &handler{cfg: cfg}

	cmds := map[string]cliCommand{
		"exit": {name: "exit", description: "Exit the pokedex", callback: handler.commandExit},
		"map":  {name: "map", description: "Displays 20 location area names. Each subsequent show the next 20 location ares", callback: handler.commandMap},
		"mapb": {name: "mapb", description: "Displays the previous 20 location area names. Each subsequent show the next previous 20 location ares", callback: handler.commandMapb},
	}

	cmds["help"] = cliCommand{name: "help", description: "Displayes the help message", callback: func() error { return handler.commandHelp(cmds) }}

	return cmds
}

func (h *handler) commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func (h *handler) commandHelp(commands map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)

	}
	return nil
}

func (h *handler) commandMap() error {
	areaUrl := h.cfg.next
	previousUrl := h.cfg.previous

	// if its the first page
	if len(areaUrl) == 0 && len(previousUrl) == 0 {
		areaUrl = "https://pokeapi.co/api/v2/location-area/"
	}

	result, err := pokedex.RetrieveLocationItems(areaUrl)

	if err != nil {
		return err
	}

	printLocationAreas(h.cfg, result)
	return nil
}

func (h *handler) commandMapb() error {
	areaUrl := h.cfg.previous

	if len(areaUrl) == 0 {
		fmt.Println("Nothing to retrieve")
		return nil
	}

	result, err := pokedex.RetrieveLocationItems(areaUrl)

	if err != nil {
		return err
	}

	printLocationAreas(h.cfg, result)
	return nil
}

func printLocationAreas(config *config, data pokedex.LocationAreaResult) {

	if data.Next != nil {
		config.next = *data.Next
	} else {
		config.next = ""
	}

	if data.Previous != nil {
		config.previous = *data.Previous
	} else {
		config.previous = ""
	}

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

}
