package pokeapi


import(
        "fmt"
        "io"
        "net/http"
        "encoding/json"
)


//call the Pokemon API with details


func(c *Client) PokemonDetailGet (pokemonName string) (PokemonDetail, error) {

        url := baseURL + "/pokemon/" + pokemonName

        check, ok := c.cache.Get(url)
        if ok == true {
                var pokeDetail PokemonDetail
                if err := json.Unmarshal(check, &pokeDetail); err != nil {
                        return PokemonDetail{}, err
                } else {
                        return pokeDetail, nil
                }
        }

	res, err := http.Get(url)
        if err != nil {
                return PokemonDetail{}, fmt.Errorf("error creating request: %w", err)
        }


        body, err := io.ReadAll(res.Body)
        if err != nil {
                return PokemonDetail{}, fmt.Errorf("reading error: %w", err)
        }
        defer res.Body.Close()


        if res.StatusCode > 299 {
                return PokemonDetail{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)

        }

        var pokeDetail PokemonDetail
        if err := json.Unmarshal(body, &pokeDetail); err != nil {
                return PokemonDetail{}, err

        }

        c.cache.Add(url, body)
        return pokeDetail, nil

}
