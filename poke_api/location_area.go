package pokeapi

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type locationAreasResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationAreas struct {
	CurrentPageIndex uint
	LastPageIndex    uint
	Areas            []LocationArea
}
