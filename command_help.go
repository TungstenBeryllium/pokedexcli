package main

import(
	"fmt"

)



func commandHelp(cfg *config, args ...string) error {

        message := `Welcome to the Pokedex!
Usage:

explore area-name: Lists Pokemon in the area
help: Displays a help message
exit: Exit the Pokedex`

        fmt.Println(message)
        return nil
}
