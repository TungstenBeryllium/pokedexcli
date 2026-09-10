package pokeapi

import (

	"time"
	"github.com/TungstenBeryllium/pokedexcli/internal/pokecache"

)



type Client struct {
	cache *pokecache.Cache
}


func NewClient() (*Client) {

	Newcache := pokecache.NewCache(5*time.Second)

	NClient := Client{
		cache: Newcache,
	}

	return  &NClient

}
