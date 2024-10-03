package pokeapi

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"

	pokecache "github.com/1DIce/pokedexcli/pokecache"
)

var locationAreasCache *pokecache.Cache[locationAreasResponse]

func getLocationsCache() *pokecache.Cache[locationAreasResponse] {
	if locationAreasCache == nil {
		newCache := pokecache.NewCache[locationAreasResponse](30 * time.Second)
		locationAreasCache = &newCache
	}
	return locationAreasCache
}

func GetLocationAreas(index uint) (LocationAreas, error) {
	url := getAreasUrl(index)

	cachedResponse, ok := getLocationsCache().Get(url)
	if ok && cachedResponse != nil {
		return toLocationAreas(*cachedResponse, index), nil
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("network error: %v", err)
	}

	if res.StatusCode > 299 {
		return LocationAreas{}, fmt.Errorf("requesting location areas failed with %s", res.Status)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	var data locationAreasResponse
	if err := decoder.Decode(&data); err != nil {
		return LocationAreas{}, err
	}

	getLocationsCache().Set(url, data)

	return toLocationAreas(data, index), nil
}

func getAreasUrl(index uint) string {
	offset := index * 20
	return fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=20", offset)
}

func toLocationAreas(response locationAreasResponse, pageIndex uint) LocationAreas {
	lastPage := math.Ceil(float64(response.Count)/20.0) - 1

	return LocationAreas{
		CurrentPageIndex: pageIndex,
		LastPageIndex:    uint(lastPage),
		Areas:            response.Results,
	}
}
