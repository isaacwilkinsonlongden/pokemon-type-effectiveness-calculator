package main

import (
	"fmt"
	"sort"
)

func commandPokemon(cfg *config, words ...string) error {
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(words[0])
	if err != nil {
		return err
	}

	pokemonTypes := []string{}
	for _, pokemonType := range pokemonResp.Types {
		pokemonTypes = append(pokemonTypes, pokemonType.Type.Name)
	}

	multipliers := map[string]float64{}

	allTypes := []string{
		"normal", "fire", "water", "electric", "grass", "ice",
		"fighting", "poison", "ground", "flying", "psychic",
		"bug", "rock", "ghost", "dragon", "dark", "steel", "fairy",
	}

	for _, t := range allTypes {
		multipliers[t] = 1.0
	}

	for _, pokemonType := range pokemonTypes {
		typeResp, err := cfg.pokeapiClient.GetType(pokemonType)
		if err != nil {
			return err
		}

		for _, t := range typeResp.DamageRelations.DoubleDamageFrom {
			multipliers[t.Name] *= 2
		}

		for _, t := range typeResp.DamageRelations.HalfDamageFrom {
			multipliers[t.Name] *= 0.5
		}

		for _, t := range typeResp.DamageRelations.NoDamageFrom {
			multipliers[t.Name] *= 0
		}
	}

	fmt.Printf("Name: %s\n", pokemonResp.Name)
	fmt.Printf("Types: %v\n", pokemonTypes)
	printTypeEffectiveness("WEAKNESSES", multipliers)
	printTypeEffectiveness("RESISTANCES", multipliers)
	printTypeEffectiveness("IMMUNITIES", multipliers)
	printTypeEffectiveness("NORMAL", multipliers)

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
