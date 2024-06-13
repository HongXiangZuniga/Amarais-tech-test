package pokemon

type PokemonRepo interface {
	SavePokemon() error
}
