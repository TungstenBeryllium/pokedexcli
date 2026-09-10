package main

import (
        "github.com/TungstenBeryllium/pokedexcli/internal/pokeapi"
        "fmt"
	"errors"
)


func printPokemon(pokemon []pokeapi.PokemonEncounter) {


        for _, poke :=  range pokemon  {

		name := poke.Pokemon.Name
		fmt.Printf(" - %s\n", name)
        }
}


func commandExplore(cfg *config, args ...string) error {
	

	if len(args) != 1{
		return errors.New("wrong amount of arguments")

	}

	areaName := args[0]

	resp, err := cfg.client.EncounterGet(areaName)
        if err != nil {
                return err
        }

	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Printf("Found Pokemon:\n")
	printPokemon(resp)

        return nil

}


//modify for exploration
//needs to call the command from pokeapi to perform the get requesst
