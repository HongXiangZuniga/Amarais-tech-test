package pokemon

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

type port struct {
	PokeApiUrl  string
	PokemonRepo PokemonRepo
	httpClient  http.Client
}

type PokemonService interface {
	GetPokemon(id string) (*Pokemon, error)
	SavePokemon(pokemon Pokemon) error
}

func NewService(repo PokemonRepo, client http.Client, url string) PokemonService {
	return &port{
		PokemonRepo: repo,
		httpClient:  client,
		PokeApiUrl:  url,
	}
}

func (port *port) GetPokemon(id string) (*Pokemon, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	isExist, err := port.PokemonRepo.IsExist(idInt)
	if err != nil {
		return nil, err
	}
	if *isExist {
		pokemon, err := port.PokemonRepo.GetPokemon(idInt)
		if err != nil {
			return nil, err
		}
		return pokemon, nil
	}
	resp, err := port.httpClient.Get(port.PokeApiUrl + "v2/pokemon/" + id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return nil, err
	}
	pokemon.Id = idInt
	err = port.PokemonRepo.SavePokemon(pokemon)
	if err != nil {
		return nil, err
	}
	return &pokemon, nil
}

func (port *port) SavePokemon(pokemon Pokemon) error {
	isExist, err := port.PokemonRepo.IsExist(pokemon.Id)
	if err != nil {
		return err
	}
	if *isExist {
		return errors.New("Pokemon already exists")
	}
	err = port.PokemonRepo.SavePokemon(pokemon)
	if err != nil {
		return err
	}
	return nil
}
