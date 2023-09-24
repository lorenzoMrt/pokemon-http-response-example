package pokeapi

import (
	pokemoncli 
)

const (
	pokemonEndpoint = "/pokemon"
	pokemonUrl      = "https://pokeapi.co/api/v2"
)

type pokemonRepo struct {
	url string
}

func NewPokemonRepository(url string) *pokemonRepo {
	return &pokemonRepo{
		url: url,
	}
}

func (p *pokemonRepo) GetPokemon() (*Pokemon, error) {
	pokemon := &Pokemon{}
	err := p.Get(pokemonUrl+pokemonEndpoint, pokemon)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}