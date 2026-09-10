package main

import (
	"strings"
	"os"
	"bufio"
	"fmt"
        "github.com/TungstenBeryllium/pokedexcli/internal/pokeapi"
	

)


func cleanInput(text string) []string {
		
	lowerStr := strings.ToLower(text)
	cleaned := strings.Fields(lowerStr)
	return cleaned
}


type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
            	"exit": {
                        name:        "exit",
                        description: "Exit the Pokedex",
                        callback:    commandExit,
                },

                "help": {
                        name: "help",
                        description: "Help menu",
                        callback: commandHelp,
                },

		"map": {
			name: "map",
			description: "Call map or move it forward",
			callback: commandMap,
		},

		"mapb": {
			name: "mapb",
			description: "Return to the previous page of the map",
			callback: commandMapb,
		},

		"explore": {
			name: "explore",
			description: "Explore the area for Pokemon",
			callback: commandExplore,

		},

		"catch": {
			name: "catch",
			description: "Catch Pokemon",
			callback: commandCatch,

		},

		"inspect": {
			name: "inspect",
			description: "Show Pokemon data",
			callback: commandInspect,
		},

		"pokedex": {
			name: "pokedex",
			description: "Show caught Pokemon",
			callback: commandPokedex,
		},
        }
}


type config struct {

	registry map[string]cliCommand
	
	client *pokeapi.Client

	next *string

	previous *string

	caughtPokemon map[string]pokeapi.PokemonDetail
}


func repl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)	
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)

		if len(cleanedInput) == 0 {
			continue
		}


		commandName := cleanedInput[0]
		arguments := cleanedInput[1:]

		val, ok := cfg.registry[commandName]



		if ok {
			err := val.callback(cfg, arguments...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}

}

	
