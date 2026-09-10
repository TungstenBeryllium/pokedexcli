package main

import(

	"fmt"

)


func commandPokedex(cfg *config, args ...string) error {

	pokedex := cfg.caughtPokemon

	fmt.Println("Your Pokedex:")
		
		for _, pkmn := range pokedex {

			fmt.Printf("	- %s\n", pkmn.PokemonName)

		}
		

        return nil
}
