package pokemon

type PokemonRepo interface {
	SavePokemon(Pokemon) error
	IsExist(id int) (*bool, error)
	GetPokemon(id int) (*Pokemon, error)
}
