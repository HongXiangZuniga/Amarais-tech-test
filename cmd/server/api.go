package main

import (
	"net/http"

	"github.com/HongXiangZuniga/Amarais-tech-test/pkg/adapter/persistence/mongodb"
	"github.com/HongXiangZuniga/Amarais-tech-test/pkg/pokemon"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Starting server...")
	httpClient := http.Client{}
	pokemonRepo := mongodb.NewPokemonRepo()
	pokemonService := pokemon.NewService(pokemonRepo, httpClient)
	pokemonService.GetPokemon(1)
}
