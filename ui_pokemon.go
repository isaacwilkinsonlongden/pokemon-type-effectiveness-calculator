package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/isaacwilkinsonlongden/pokemon-weakness-calculator/internal/effectiveness"
)

func printPokemonResult(result effectiveness.Result) {
	fmt.Println()
	fmt.Printf("Pokemon: %s\n", result.Name)
	fmt.Printf("Types: %s\n", formatTypes(result.Types))
	fmt.Println()

	printEffectivenessSection("Weaknesses", result.Multipliers, func(v float64) bool { return v > 1 })
	printEffectivenessSection("Resistances", result.Multipliers, func(v float64) bool { return v < 1 && v > 0 })
	printEffectivenessSection("Immunities", result.Multipliers, func(v float64) bool { return v == 0 })
	printEffectivenessSection("Normal", result.Multipliers, func(v float64) bool { return v == 1 })
}

func printEffectivenessSection(
	title string,
	multipliers map[string]float64,
	include func(float64) bool,
) {
	fmt.Println(sectionTitle(title))

	keys := make([]string, 0, len(multipliers))
	for k, v := range multipliers {
		if include(v) {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	if len(keys) == 0 {
		fmt.Println("  none")
		fmt.Println()
		return
	}

	for _, k := range keys {
		fmt.Printf("  %-18s x%.1f\n", colorizeType(k), multipliers[k])
	}
	fmt.Println()
}

func formatTypes(types []string) string {
	if len(types) == 0 {
		return "[]"
	}

	colored := make([]string, 0, len(types))
	for _, t := range types {
		colored = append(colored, colorizeType(t))
	}

	return strings.Join(colored, ", ")
}

func sectionTitle(title string) string {
	return styleBold(title + ":")
}
