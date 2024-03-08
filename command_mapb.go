package main

import (
	"fmt"

	pokeapi "github.com/0xRTH/pokedex/internal/pokeApi"
)

func mapb(config *config, args ...string) error {
	url := config.previousMap
	if url == "" {
		fmt.Println("No previous maps")
		return nil
	}
	apiMap, _ := pokeapi.GetMap(url)

	for _, area := range apiMap.Results {
		fmt.Println(area.Name)
	}
	config.nextMap = apiMap.Next
	config.previousMap = apiMap.Previous
	return nil
}
