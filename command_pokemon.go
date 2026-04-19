package main

import (
	"github.com/isaacwilkinsonlongden/pokemon-weakness-calculator/internal/effectiveness"
)

func commandPokemon(cfg *config, words ...string) error {
	result, err := effectiveness.Calculate(cfg.pokeapiClient, words[0], cfg.generation)
	if err != nil {
		return err
	}

	printPokemonResult(result)
	return nil
}
