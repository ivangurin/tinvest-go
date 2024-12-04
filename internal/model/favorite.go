package model

type Favorite struct {
	ID     string
	Ticker string
	Kind   string
}

type Favorites []*Favorite
