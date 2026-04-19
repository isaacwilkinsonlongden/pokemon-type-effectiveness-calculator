package main

import (
	"fmt"
	"sort"

	"github.com/isaacwilkinsonlongden/pokemon-weakness-calculator/internal/effectiveness"
)

func commandPokemon(cfg *config, words ...string) error {
	result, err := effectiveness.Calculate(cfg.pokeapiClient, words[0], cfg.generation)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n", result.Name)
	fmt.Printf("Types: %v\n", result.Types)
	printTypeEffectiveness("WEAKNESSES", result.Multipliers)
	printTypeEffectiveness("RESISTANCES", result.Multipliers)
	printTypeEffectiveness("IMMUNITIES", result.Multipliers)
	printTypeEffectiveness("NORMAL", result.Multipliers)

	return nil
}

func printTypeEffectiveness(damageType string, multipliers map[string]float64) {
	fmt.Printf("%s:\n", damageType)

	keys := make([]string, 0, len(multipliers))
	switch damageType {
	case "WEAKNESSES":
		for k, v := range multipliers {
			if v > 1 {
				keys = append(keys, k)
			}
		}
	case "RESISTANCES":
		for k, v := range multipliers {
			if v < 1 && v > 0 {
				keys = append(keys, k)
			}
		}
	case "IMMUNITIES":
		for k, v := range multipliers {
			if v == 0 {
				keys = append(keys, k)
			}
		}
	case "NORMAL":
		for k, v := range multipliers {
			if v == 1 {
				keys = append(keys, k)
			}
		}
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("- %s (x%.1f) ", k, multipliers[k])
	}
	fmt.Println()
}
