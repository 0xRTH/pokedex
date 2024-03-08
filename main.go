package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/0xRTH/pokedex/internal/pokeApi"
	"github.com/0xRTH/pokedex/internal/pokecache"
)

type userInput struct {
	command string
	args    []string
}

type config struct {
	nextMap     string
	previousMap string
	cache       pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "Help",
			description: "Show this help message",
			callback:    help,
		},
		"exit": {
			name:        "Exit",
			description: "Exit the program",
			callback:    exit,
		},
		"map": {
			name:        "Map",
			description: "Get next 20 area",
			callback:    getMap,
		},
		"mapb": {
			name:        "Mapb",
			description: "Get previous 20 area",
			callback:    mapb,
		},
		"explore": {
			name:        "Explore",
			description: "Explore a given area by entering its name",
			callback:    explore,
		},
		"catch": {
			name:        "Catch",
			description: "Try to catch a pokemon by entering its name",
			callback:    catch,
		},
		"inspect": {
			name:        "Inspect",
			description: "Inspect pokemon if he is in your pokedex",
			callback:    inspect,
		},
		"pokedex": {
			name:        "Pokedex",
			description: "List all your captured pokemon available in your pokedex",
			callback:    pokedex,
		},
	}
}

var capturedPokemons = map[string]pokeapi.PokemonInfo{}

func main() {
	reader := bufio.NewScanner(os.Stdin)

	config := config{
		nextMap:     "https://pokeapi.co/api/v2/location-area",
		previousMap: "",
	}

	input := userInput{}

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := strings.Fields(reader.Text())
		input.command = words[0]
		input.args = words[1:]

		if command, ok := getCommands()[input.command]; ok {
			err := command.callback(&config, input.args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Command not found...")
			continue
		}
	}
}
