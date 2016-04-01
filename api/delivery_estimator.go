package api

import (
	"time"

	"github.com/tjarratt/go-best-practices/domain"
)

//go:generate counterfeiter . PizzaDeliveryEstimator
type PizzaDeliveryEstimator interface {
	EstimatedDeliveryTime(domain.Pizza) time.Duration
}

func NewPizzaDeliveryEstimator() PizzaDeliveryEstimator {
	return pizzaDeliveryEstimator{}
}

type pizzaDeliveryEstimator struct{}

func (estimator pizzaDeliveryEstimator) EstimatedDeliveryTime(pizza domain.Pizza) time.Duration {
	return 30 * time.Minute // or less!
}
