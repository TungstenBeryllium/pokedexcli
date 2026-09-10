package main

import (
	"github.com/TungstenBeryllium/pokedexcli/internal/pokeapi"
	"fmt"
	"errors"
)


func printOut(areas []pokeapi.LocationArea) {


	for _, location :=  range areas {

		fmt.Println(location.Name)			
	}
}



func commandMap(cfg *config, args ...string) error {
	
	response, err := cfg.client.LocationGet(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = response.Next
	cfg.previous = response.Previous
	printOut(response.Results)

        return nil

}


func commandMapb(cfg *config, args ...string) error {


        if cfg.previous == nil {
               return errors.New("you're on the first page")
	}	

        response, err  := cfg.client.LocationGet(cfg.previous)
	if err != nil {
                return err
        }
	
	cfg.next = response.Next
        cfg.previous = response.Previous
        printOut(response.Results)

        return nil

}

