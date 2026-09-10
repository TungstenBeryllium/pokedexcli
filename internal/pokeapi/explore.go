package pokeapi


import(
	"fmt"
        "io"
        "net/http"
        "encoding/json"
)


func(c *Client) EncounterGet (areaName string) ([]PokemonEncounter, error) {

	url := baseURL + "/location/" + areaName

	check, ok := c.cache.Get(url)
        if ok == true {
                var areaDetail LocationAreaDetail
                if err := json.Unmarshal(check, &areaDetail); err != nil {
                        return nil, err
                } else {
                        return areaDetail.Encounters, nil
                }
        }

	res, err := http.Get(url)
        if err != nil {
                return nil, fmt.Errorf("error creating request: %w", err)
        }


        body, err := io.ReadAll(res.Body)
        if err != nil {
                return nil, fmt.Errorf("reading error: %w", err)
        }
        defer res.Body.Close()


        if res.StatusCode > 299 {
                return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)

        }

        var areaDetail LocationAreaDetail
        if err := json.Unmarshal(body, &areaDetail); err != nil {
                return nil, err

        }

        c.cache.Add(url, body)
        return areaDetail.Encounters,  nil

}

