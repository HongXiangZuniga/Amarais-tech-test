package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/HongXiangZuniga/Amarais-tech-test/pkg/adapter/http/rest"
	"github.com/HongXiangZuniga/Amarais-tech-test/pkg/adapter/persistence/mongodb"
	"github.com/HongXiangZuniga/Amarais-tech-test/pkg/pokemon"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Starting server...")
	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatal("Failed to load .env file", zap.Error(err))
	}
	//Load env
	url := os.Getenv("POKE_API_URL")
	token := os.Getenv("API_TOKEN")
	httpClient := http.Client{}
	//Init Repo and Services
	mdbURI := os.Getenv("MONGO_URI")
	mdbName := os.Getenv("MONGO_DBNAME")
	mdb := initMongo(mdbURI, mdbName)
	pokemonRepo := mongodb.NewPokemonRepo(mdb)
	pokemonService := pokemon.NewService(pokemonRepo, httpClient, url)
	//Init Handler y router
	pokemonHandler := rest.NewPokemonHandler(pokemonService)
	router := rest.NewRouter(pokemonHandler, token)
	router.Run(":8080")
}

func initMongo(uri string, dbName string) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("Not mongo connection, error: " + err.Error())
		os.Exit(0)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	err = client.Connect(ctx)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println("Not mongo connection, error: " + err.Error())
		os.Exit(0)
	}
	return client.Database(dbName)
}
