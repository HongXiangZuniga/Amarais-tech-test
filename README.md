# Amarais-tech-test


#### Uso
Para Ejecutar

```
Make run
or
docker-compose up -d
```

#### Diagramas
![Link to Amarais-tech-test](/desing.png)

### Estructura de directorio

```
.
├── Api.postman_collection.json
├── Makefile
├── README.md
├── build
│   └── Dockerfile
├── cmd
│   └── server
│       └── api.go
├── db
├── diagrama.png
├── docker-compose.yml
├── go.mod
├── go.sum
└── pkg
    ├── adapter
    │   ├── http
    │   │   └── rest
    │   │       ├── pokemon.go
    │   │       └── router.go
    │   └── persistence
    │       └── mongodb
    │           └── pokemon.go
    └── pokemon
        ├── model.go
        ├── repository.go
        └── service.go
```

#### Endpoints
```
GET    /pokemon/:id           
POST   /pokemon/    
```
Disponible postman Collection para test.

#### Env de ejemplo
```
POKE_API_URL=https://pokeapi.co/api/
API_TOKEN="1234"
MONGO_URI=mongodb://mongo:27017
MONGO_DBNAME=pokemon
MONGO_COLLECTION=pokemon
```
