package pokemon

import (
	"fmt"
	"net/http"
)

type port struct {
	PokemonRepo PokemonRepo
	httpClient  http.Client
}

type PokemonService interface {
	GetPokemon(id int) error
	SavePokemon() error
}

func NewService(repo PokemonRepo, client http.Client) PokemonService {
	return &port{
		PokemonRepo: repo,
		httpClient:  client,
	}
}

func (p *port) GetPokemon(id int) error {
	fmt.Println(id)
	return nil
}

func (p *port) SavePokemon() error {
	return nil
}
