package mongodb

import "github.com/HongXiangZuniga/Amarais-tech-test/pkg/pokemon"

type port struct {
}

func NewPokemonRepo() pokemon.PokemonRepo {
	return &port{}
}

func (port *port) SavePokemon() error {
	return nil
}
