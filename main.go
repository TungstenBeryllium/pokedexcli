package main

import (
        "github.com/TungstenBeryllium/pokedexcli/internal/pokeapi"
        
)

func main() {


	myLocalRegistry := getCommands()
	
	currentPokedex := map[string]pokeapi.PokemonDetail{}

	client := pokeapi.NewClient()

	cfg := &config{
		registry: myLocalRegistry,
		client: client,
		caughtPokemon: currentPokedex,

}

	repl(cfg)
}
