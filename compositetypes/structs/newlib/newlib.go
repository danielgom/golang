package newlib

type Account struct {
	Name    string
	Email   string
	Balance float64
}

type Private struct {
	secret string
	value  string
}

type Cellphone struct {
	brand string
	size int
	screen float64
}

func NewCellphone() Cellphone {
	return Cellphone{"Iphone", 10, 5.8}
}

