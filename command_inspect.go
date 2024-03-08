package main

import (
	"errors"
	"fmt"
)

func inspect(config *config, args ...string) error {
	if len(args) > 1 {
		return errors.New("too many pokemons")
	} else if len(args) < 1 {
		return errors.New("please provide pokemon name")
	}

	if pokemon, ok := capturedPokemons[args[0]]; ok {
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  - %s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pokemonType := range pokemon.Types {
			fmt.Printf("  - %s\n", pokemonType.Type.Name)
		}
	} else {
		fmt.Println("You don't have captured this pokemon yet")
	}
	return nil
}
