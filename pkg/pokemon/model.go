package pokemon

type Pokemon struct {
	Name      string        `json:"name"`
	Order     int           `json:"order"`
	Abilities []abilitie    `json:"abilities"`
	Types     []typePokemon `json:"types"`
}

type abilitie struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type typePokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
