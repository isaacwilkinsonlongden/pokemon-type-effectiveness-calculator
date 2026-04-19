package effectiveness

import "github.com/isaacwilkinsonlongden/pokemon-weakness-calculator/internal/pokeapi"

type Result struct {
	Name        string
	Types       []string
	Multipliers map[string]float64
}

func AllTypesForGeneration(gen Generation) []string {
	switch gen {
	case GenerationIII:
		return []string{
			"normal", "fire", "water", "electric", "grass", "ice",
			"fighting", "poison", "ground", "flying", "psychic",
			"bug", "rock", "ghost", "dragon", "dark", "steel",
		}
	default:
		return []string{
			"normal", "fire", "water", "electric", "grass", "ice",
			"fighting", "poison", "ground", "flying", "psychic",
			"bug", "rock", "ghost", "dragon", "dark", "steel", "fairy",
		}
	}
}

func Calculate(client pokeapi.Client, pokemonName string, gen Generation) (Result, error) {
	pokemonResp, err := client.GetPokemon(pokemonName)
	if err != nil {
		return Result{}, err
	}

	pokemonTypes := []string{}
	for _, pokemonType := range pokemonResp.Types {
		pokemonTypes = append(pokemonTypes, pokemonType.Type.Name)
	}

	multipliers := map[string]float64{}
	for _, t := range AllTypesForGeneration(gen) {
		multipliers[t] = 1.0
	}

	for _, pokemonType := range pokemonTypes {
		typeResp, err := client.GetType(pokemonType)
		if err != nil {
			return Result{}, err
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

	return Result{
		Name:        pokemonResp.Name,
		Types:       pokemonTypes,
		Multipliers: multipliers,
	}, nil
}
