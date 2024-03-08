package main

import (
	"errors"
	"fmt"

	pokeapi "github.com/0xRTH/pokedex/internal/pokeApi"
)

func explore(config *config, args ...string) error {
	if len(args) > 1 {
		return errors.New("too many location")
	} else if len(args) < 1 {
		return errors.New("please provide location")
	}
	pokemons, _ := pokeapi.GetPokemonsInArea(args[0])
	fmt.Println("Found pokemons:")

	for _, pokemon := range pokemons {
		fmt.Println(" -", pokemon.Name)
	}
	return nil
}
