package main

import (
	"time"

	"github.com/isaacwilkinsonlongden/pokemon-weakness-calculator/internal/effectiveness"
	"github.com/isaacwilkinsonlongden/pokemon-weakness-calculator/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		generation:    effectiveness.GenerationIII,
	}

	startRepl(cfg)
}
