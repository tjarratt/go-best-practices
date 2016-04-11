package usecases

import "github.com/tjarratt/go-best-practices/domain"

type OrderPizzaRequest struct {
	Whom     string
	Address  string
	Dough    domain.DoughType
	Toppings []domain.Ingredient
}
