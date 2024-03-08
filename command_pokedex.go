package main

import "fmt"

func pokedex(config *config, args ...string) error {
	fmt.Println("Pokedex:")
	for _, pokemon := range capturedPokemons {
		fmt.Println(" -", pokemon.Name)
	}
	return nil
}
