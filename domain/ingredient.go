package domain

type Ingredient interface {
	Name() string
}

type Pepperoni struct{}

func (p Pepperoni) Name() string {
	return "pepperoni"
}
