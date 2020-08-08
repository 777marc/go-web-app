package models

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}
