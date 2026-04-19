package effectiveness

import "github.com/isaacwilkinsonlongden/pokemon-weakness-calculator/internal/pokeapi"

type Result struct {
	Name        string
	Types       []string
	Multipliers map[string]float64
}

type DamageRelations struct {
	DoubleDamageFrom []string
	HalfDamageFrom   []string
	NoDamageFrom     []string
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

	pokemonTypes := pokemonTypesForGeneration(pokemonResp, gen)

	multipliers := map[string]float64{}
	for _, t := range AllTypesForGeneration(gen) {
		multipliers[t] = 1.0
	}

	for _, pokemonType := range pokemonTypes {
		typeResp, err := client.GetType(pokemonType)
		if err != nil {
			return Result{}, err
		}

		relations := damageRelationsForGeneration(typeResp, gen)

		for _, typeName := range relations.DoubleDamageFrom {
			if _, ok := multipliers[typeName]; ok {
				multipliers[typeName] *= 2
			}
		}

		for _, typeName := range relations.HalfDamageFrom {
			if _, ok := multipliers[typeName]; ok {
				multipliers[typeName] *= 0.5
			}
		}

		for _, typeName := range relations.NoDamageFrom {
			if _, ok := multipliers[typeName]; ok {
				multipliers[typeName] *= 0
			}
		}
	}

	return Result{
		Name:        pokemonResp.Name,
		Types:       pokemonTypes,
		Multipliers: multipliers,
	}, nil
}

func damageRelationsForGeneration(t pokeapi.Type, gen Generation) DamageRelations {
	if gen == Current {
		return currentDamageRelations(t)
	}

	for _, past := range t.PastDamageRelations {
		if past.Generation.Name == string(gen) {
			return DamageRelations{
				DoubleDamageFrom: extractNames(past.DamageRelations.DoubleDamageFrom),
				HalfDamageFrom:   extractNames(past.DamageRelations.HalfDamageFrom),
				NoDamageFrom:     extractNames(past.DamageRelations.NoDamageFrom),
			}
		}
	}

	return currentDamageRelations(t)
}

func extractNames(items []struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}) []string {
	names := make([]string, 0, len(items))
	for _, item := range items {
		names = append(names, item.Name)
	}
	return names
}

func currentDamageRelations(t pokeapi.Type) DamageRelations {
	return DamageRelations{
		DoubleDamageFrom: extractNames(t.DamageRelations.DoubleDamageFrom),
		HalfDamageFrom:   extractNames(t.DamageRelations.HalfDamageFrom),
		NoDamageFrom:     extractNames(t.DamageRelations.NoDamageFrom),
	}
}

func pokemonTypesForGeneration(p pokeapi.Pokemon, gen Generation) []string {
	if gen == Current {
		return extractPokemonTypeNames(p.Types)
	}

	requestedGen := generationNumber(gen)

	for _, past := range p.PastTypes {
		pastGen := generationNumber(Generation(past.Generation.Name))

		if requestedGen <= pastGen {
			return extractPokemonTypeNames(past.Types)
		}
	}

	return extractPokemonTypeNames(p.Types)
}

func extractPokemonTypeNames(types []struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}) []string {
	names := make([]string, 0, len(types))
	for _, t := range types {
		names = append(names, t.Type.Name)
	}
	return names
}
