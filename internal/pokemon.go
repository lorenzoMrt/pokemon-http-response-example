package pokemoncli

type Pokemon struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokemonRepository interface {
	GetPokemon() (Pokemon, error)
}

func NewPokemon(count int, next string, previous string, results []struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}) (p Pokemon) {
	p = Pokemon{
		Count:    count,
		Next:     next,
		Previous: previous,
		Results:  results,
	}
	return
}
