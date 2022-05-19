package model

type Shoe struct {
	Brand string
}

func NewShoe(brand string) *Shoe {
	return &Shoe{brand}
}
