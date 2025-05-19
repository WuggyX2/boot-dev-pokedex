package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/WuggyX2/boot-dev-pokedex/internal/pokecache"
	"github.com/WuggyX2/boot-dev-pokedex/pokedex"
)

type handler struct {
	cfg      *config
	cache    *pokecache.Cache
	pokemons map[string]pokedex.PokemonResult
}

type cliCommand struct {
	name        string
	description string
	callback    func([]string) error
}

const LOCATION_AREA_URL = "https://pokeapi.co/api/v2/location-area/"

func registerCmds(cfg *config) map[string]cliCommand {
	interval := 5 * time.Second

	c := pokecache.NewCache(interval)

	handler := &handler{cfg: cfg, cache: &c, pokemons: make(map[string]pokedex.PokemonResult)}

	cmds := map[string]cliCommand{
		"exit": {name: "exit", description: "Exit the pokedex", callback: handler.commandExit},
		"map": {
			name:        "map",
			description: "Displays 20 location area names. Each subsequent show the next 20 location ares",
			callback:    handler.commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location area names. Each subsequent show the next previous 20 location ares",
			callback:    handler.commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explores the given area. Reveals the pokemon the given area. Usage 'explore {area name}'",
			callback:    handler.commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Used this command to try to catch a pokemon. Usage 'catch {pokemon name}'",
			callback:    handler.commandCatch,
		},
	}

	cmds["help"] = cliCommand{
		name:        "help",
		description: "Displayes the help message",
		callback:    func(args []string) error { return handler.commandHelp(cmds) },
	}

	return cmds
}

func (h *handler) commandExit(args []string) error {
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

func (h *handler) commandMap(args []string) error {
	areaUrl := h.cfg.next
	previousUrl := h.cfg.previous

	// if its the first page
	if len(areaUrl) == 0 && len(previousUrl) == 0 {
		areaUrl = LOCATION_AREA_URL
	}

	result, err := pokedex.RetrieveLocationItems(areaUrl, h.cache)

	if err != nil {
		return err
	}

	printLocationAreas(h.cfg, result)
	return nil
}

func (h *handler) commandMapb(args []string) error {
	areaUrl := h.cfg.previous

	if len(areaUrl) == 0 {
		fmt.Println("Nothing to retrieve")
		return nil
	}

	result, err := pokedex.RetrieveLocationItems(areaUrl, h.cache)

	if err != nil {
		return err
	}

	printLocationAreas(h.cfg, result)
	return nil
}

func (h *handler) commandExplore(args []string) error {
	if len(args) == 0 {
		return errors.New("Area name was not provided!")
	}

	areaName := args[0]
	fmt.Printf("Exploring %s...\n", areaName)

	result, err := pokedex.GetPokemonsInArea(LOCATION_AREA_URL+areaName, h.cache)

	if err != nil {
		return err
	}

	printPokemons(result)

	return nil
}

func (h *handler) commandCatch(args []string) error {
	if len(args) == 0 {
		return errors.New("Pokemon name not provied")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := pokedex.GetPokemon("https://pokeapi.co/api/v2/pokemon/"+pokemonName, h.cache)

	if err != nil {
		return err
	}

	catchRate := rand.Intn(256)

	if pokemon.BaseExperience > catchRate {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	h.pokemons[pokemon.Name] = pokemon

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

func printPokemons(data pokedex.PokemonsInAreaResult) {
	fmt.Println("Found Pokemon:")

	for _, pokemon := range data.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
}
