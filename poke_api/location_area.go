package pokeapi

type locationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type locationAreasResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []locationArea `json:"results"`
}

type LocationAreaInfo struct {
	Name string
	Id   int
}

type LocationAreas struct {
	CurrentPageIndex uint
	LastPageIndex    uint
	Areas            []LocationAreaInfo
}
