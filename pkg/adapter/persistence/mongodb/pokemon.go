package mongodb

import (
	"context"

	"github.com/HongXiangZuniga/Amarais-tech-test/pkg/pokemon"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type port struct {
	db *mongo.Database
}

func NewPokemonRepo(db *mongo.Database) pokemon.PokemonRepo {
	return &port{db: db}
}

func (port *port) SavePokemon(pokemon pokemon.Pokemon) error {
	ctx := context.Background()
	_, err := port.db.Collection("pokemons").InsertOne(ctx, pokemon)
	if err != nil {
		return err
	}
	return nil
}

func (port *port) IsExist(id int) (*bool, error) {
	isExist := false
	ctx := context.Background()
	result := port.db.Collection("pokemons").FindOne(ctx, bson.M{"idPokemon": id})
	if result.Err() != nil {
		if result.Err().Error() == "mongo: no documents in result" {
			return &isExist, nil
		}
		return nil, result.Err()
	}
	isExist = true
	return &isExist, nil
}

func (port *port) GetPokemon(id int) (*pokemon.Pokemon, error) {
	ctx := context.Background()
	result := port.db.Collection("pokemons").FindOne(ctx, bson.M{"idPokemon": id})
	if result.Err() != nil {
		return nil, result.Err()
	}
	var pokemon pokemon.Pokemon
	err := result.Decode(&pokemon)
	if err != nil {
		return nil, err
	}
	return &pokemon, nil
}
