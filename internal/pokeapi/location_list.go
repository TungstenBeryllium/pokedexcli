package pokeapi



import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)



func(c *Client) LocationGet(pageURL * string) (LocationAreaResponse, error) {

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}


	check, ok := c.cache.Get(url)
	if ok == true {
		var response LocationAreaResponse
		if err := json.Unmarshal(check, &response); err != nil {
                	return LocationAreaResponse{}, err
        	} else {
			return response, nil
		}
	}


	res, err := http.Get(url)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("reading error: %w", err)
        }
	defer res.Body.Close()


	if res.StatusCode > 299 {
		return LocationAreaResponse{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)

	}
	
	var response LocationAreaResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return LocationAreaResponse{}, err
		
	}
	
	c.cache.Add(url, body)
	return response, nil


	
}
