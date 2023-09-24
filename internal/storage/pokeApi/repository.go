package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	pokemoncli "pokemon-cli/internal"
)

const (
	pokemonEndpoint = "/pokemon"
	pokemonUrl      = "https://pokeapi.co/api/v2"
)

type pokemonRepo struct {
	url string
}

func NewPokemonRepository(url string) pokemoncli.PokemonRepository {
	return &pokemonRepo{
		url: url,
	}
}

func (p *pokemonRepo) GetPokemon() (pokemon pokemoncli.Pokemon, err error) {
	response, err := http.Get(pokemonUrl + pokemonEndpoint)
	if err != nil {
		return nil, err
	}
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(contents, &pokemon)
	if err != nil {
		return nil, err
	}
	if response == nil {
		return pokemoncli.Pokemon{}, nil
	}
	return
}
