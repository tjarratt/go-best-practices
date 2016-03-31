package api

import (
	"github.com/tjarratt/go-best-practices/domain"
)

//go:generate counterfeiter . PizzaRepository
type PizzaRepository interface {
	MakePizza(domain.DoughType, []domain.Ingredient) (domain.Pizza, int64)
}
