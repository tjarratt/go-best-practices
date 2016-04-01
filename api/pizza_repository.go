package api

import (
	"github.com/tjarratt/go-best-practices/domain"
)

//go:generate counterfeiter . PizzaRepository
type PizzaRepository interface {
	MakePizza(domain.DoughType, []domain.Ingredient) (domain.Pizza, int64, error)
}

func NewPizzaRepository() PizzaRepository {
	return pizzaRepository{}
}

type pizzaRepository struct{}

func (repo pizzaRepository) MakePizza(
	dough domain.DoughType,
	toppings []domain.Ingredient,
) (domain.Pizza, int64, error) {
	return domain.Pizza{}, 0, nil
}
