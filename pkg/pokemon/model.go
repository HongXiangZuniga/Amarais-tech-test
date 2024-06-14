package pokemon

type Pokemon struct {
	Name      string      `json:"name" bson:"name"`
	Order     int         `json:"order" bson:"order"`
	Height    int         `json:"height" bson:"height"`
	Weight    int         `json:"weight" bson:"weight"`
	Abilities []abilities `json:"abilities" bson:"abilities"`
	Id        int         `json:"id" bson:"idPokemon"`
	Pokemon   string      `json:"pokemon,omitempty" bson:"pokemon,omitempty"`
}

type ability struct {
	Name string `json:"name" bson:"name"`
	Url  string `json:"url" bson:"url"`
}

type abilities struct {
	Ability ability `json:"ability" bson:"ability"`
	Hidden  bool    `json:"is_hidden" bson:"is_hidden"`
	Slot    int     `json:"slot" bson:"slot"`
}
