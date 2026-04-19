package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/isaacwilkinsonlongden/pokemon-weakness-calculator/internal/effectiveness"
	"github.com/isaacwilkinsonlongden/pokemon-weakness-calculator/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	generation    effectiveness.Generation
}

func startRepl(cfg *config) {
	err := commandHelp(cfg, "")
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Pokemon name > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, words...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

		err = commandPokemon(cfg, commandName)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    commandExit,
		},
		"pokemon": {
			name:        "pokemon",
			description: "Get pokemon weakness",
			callback:    commandPokemon,
		},
	}
}
