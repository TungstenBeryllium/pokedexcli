package main


import (
        "fmt"
        "errors"
)


func commandInspect(cfg *config, args ...string) error {


        if len(args) != 1{
                return errors.New("wrong amount of arguments")

        }

        pokeName := args[0]
	
	if caughtPokemon, ok := cfg.caughtPokemon[pokeName]; !ok {

		fmt.Printf("You have not caught %s\n", pokeName)
	
	} else {
		
		fmt.Printf("Name: %s\n", caughtPokemon.PokemonName)
		fmt.Printf("Height: %d\n", caughtPokemon.Height)
		fmt.Printf("Weight: %d\n", caughtPokemon.Weight)
		fmt.Printf("Stats:\n")

		for _, statistic := range caughtPokemon.Stats {
			
			fmt.Printf("	-%s: %d\n", statistic.Stat.StatisticName, statistic.BaseStat)

		}

		fmt.Printf("Types:\n")

		for _, typing := range caughtPokemon.Types {

			fmt.Printf("	- %s\n", typing.Type.TypeName)

		}
	}

        return nil
//optimize for inspect
}
