package rest

import (
	"github.com/HongXiangZuniga/Amarais-tech-test/pkg/pokemon"
	"github.com/gin-gonic/gin"
)

type PokemonHandler interface {
	GetPokemon(*gin.Context)
	PostPokemon(*gin.Context)
}

type port struct {
	pokemonService pokemon.PokemonService
}

func NewPokemonHandler(pokemonService pokemon.PokemonService) PokemonHandler {
	return &port{pokemonService: pokemonService}
}

func (port *port) GetPokemon(ctx *gin.Context) {
	id := ctx.Param("id")
	pokemon, err := port.pokemonService.GetPokemon(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, pokemon)
}
func (port *port) PostPokemon(ctx *gin.Context) {
	var pokemon pokemon.Pokemon
	err := ctx.BindJSON(&pokemon)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if pokemon.Name == "" {
		if pokemon.Pokemon == "" {
			ctx.JSON(400, gin.H{"error": "Name or Pokemon is required"})
		}
		pokemon.Name = pokemon.Pokemon
	}
	pokemon = *port.normalizePokemon(pokemon)
	err = port.pokemonService.SavePokemon(pokemon)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Pokemon saved"})
}

func (port *port) normalizePokemon(poke pokemon.Pokemon) *pokemon.Pokemon {
	var pokemonAux pokemon.Pokemon
	pokemonAux.Name = poke.Name
	pokemonAux.Order = poke.Order
	pokemonAux.Height = poke.Height
	pokemonAux.Weight = poke.Weight
	pokemonAux.Id = poke.Id
	pokemonAux.Abilities = poke.Abilities
	return &pokemonAux
}
