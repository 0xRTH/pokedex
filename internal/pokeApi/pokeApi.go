package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/0xRTH/pokedex/internal/pokecache"
)

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokeApiLocationDetails struct {
	PokemonEncounters []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type PokemonInfo struct {
	Name           string
	Height         int
	Weight         int
	BaseExperience int `json:"base_experience"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

type PokeApiMap struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var cache = pokecache.NewCache(10 * time.Second)

func getWithCache(url string) (body []byte) {
	if entry, ok := cache.Data[url]; ok {
		body = entry.Val
	} else {
		resp, _ := http.Get(url)
		body, _ = io.ReadAll(resp.Body)
		cache.Data[url] = pokecache.CacheEntry{
			CreatedAt: time.Now(),
			Val:       body,
		}
	}
	return body
}

func GetPokemonInfos(pokemon string) (PokemonInfo, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	body := getWithCache(url)
	infos := PokemonInfo{}
	json.Unmarshal(body, &infos)
	return infos, nil
}

func GetPokemonsInArea(area string) ([]Pokemon, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + area
	body := getWithCache(url)
	areaDetails := PokeApiLocationDetails{}
	json.Unmarshal(body, &areaDetails)
	pokemons := []Pokemon{}
	for _, pokemonEncounter := range areaDetails.PokemonEncounters {
		pokemons = append(pokemons, pokemonEncounter.Pokemon)
	}
	return pokemons, nil
}

func GetMap(url string) (PokeApiMap, error) {
	var apiMap PokeApiMap
	body := getWithCache(url)
	json.Unmarshal(body, &apiMap)
	return apiMap, nil
}
