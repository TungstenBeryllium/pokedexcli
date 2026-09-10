package main

import (
        "fmt"
        "errors"
	"math/rand"
)


func commandCatch(cfg *config, args ...string) error {


        if len(args) != 1{
                return errors.New("wrong amount of arguments")

        }

        pokeName := args[0]

        resp, err := cfg.client.PokemonDetailGet(pokeName)
        if err != nil {
                return err
        }

        fmt.Printf("Throwing a Pokeball at %s...\n", pokeName)


	if rand.Intn(700) >= resp.BaseExperience {
		fmt.Printf("%s was caught!\n", pokeName)
                cfg.caughtPokemon[resp.PokemonName] = resp

	} else {

		fmt.Printf("%s escaped!\n", pokeName)
		
	}

        return nil

}
