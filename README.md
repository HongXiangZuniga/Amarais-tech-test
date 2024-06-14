# Amarais-tech-test


#### Uso

- Crear .env
- Ejecutar con:

```
Make run
or
docker-compose up -d
```

#### Env de ejemplo
```
POKE_API_URL=https://pokeapi.co/api/
API_TOKEN="1234"
MONGO_URI=mongodb://mongo:27017
MONGO_DBNAME=users
MONGO_COLLECTION=users
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