package main

import (
	"errors"
	"fmt"
	"math/rand"

	pokeapi "github.com/0xRTH/pokedex/internal/pokeApi"
)

func catch(config *config, args ...string) error {
	if len(args) > 1 {
		return errors.New("too many pokemons")
	} else if len(args) < 1 {
		return errors.New("please provide pokemon name")
	}
	infos, _ := pokeapi.GetPokemonInfos(args[0])
	baseExp := infos.BaseExperience
	if rand.Float32()*400 > float32(baseExp) {
		fmt.Println("You catched", args[0], "!")
		if _, ok := capturedPokemons[args[0]]; !ok {
			capturedPokemons[args[0]] = infos
		}
	} else {
		fmt.Println("Too bad,", args[0], "escaped !")
	}
	return nil
}
