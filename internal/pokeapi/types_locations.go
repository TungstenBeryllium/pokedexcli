package pokeapi


type LocationArea struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

type LocationAreaResponse struct {
    Count    int            `json:"count"`
    Next     *string        `json:"next"`
    Previous *string        `json:"previous"`
    Results  []LocationArea `json:"results"`
}


type PokemonEncounter struct {
    Pokemon struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"pokemon"`
}


type LocationAreaDetail struct {
	Name string `json:"name"`
	Encounters []PokemonEncounter `json:"pokemon_encounters"`

}


type PokemonDetail struct {
	PokemonName string `json:"name"`
	BaseExperience int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []PokemonStat `json:"stats"`
	Types []PokemonType `json:"types"`
}

type PokemonStat struct {
	BaseStat int `json:"base_stat"`
	Stat Statistic `json:"stat"`
}

type Statistic struct {
	StatisticName string `json:"name"`
}

type PokemonType struct {
	Type TypeData `json:"type"`
}

type TypeData struct {
	TypeName string `json:"name"`

}
