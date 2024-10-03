package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	pokecache "github.com/1DIce/pokedexcli/pokecache"
)

type LocationResponse struct {
	ID                int                 `json:"id"`
	Name              string              `json:"name"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}

var _locationCache *pokecache.Cache[LocationResponse]

func getLocationCache() *pokecache.Cache[LocationResponse] {
	if _locationCache == nil {
		newCache := pokecache.NewCache[LocationResponse](30 * time.Second)
		_locationCache = &newCache
	}
	return _locationCache
}

func FetchLocation(nameOrId string) (LocationResponse, error) {
	url := getLocationUrl(nameOrId)

	cachedResponse, ok := getLocationCache().Get(url)
	if ok && cachedResponse != nil {
		return (*cachedResponse), nil
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("network error: %v", err)
	}

	if res.StatusCode > 299 {
		return LocationResponse{}, fmt.Errorf("requesting location areas failed with %s", res.Status)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	var data LocationResponse
	if err := decoder.Decode(&data); err != nil {
		return LocationResponse{}, err
	}

	getLocationCache().Set(url, data)

	return data, nil
}

func getLocationUrl(name string) string {
	return fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", name)
}
