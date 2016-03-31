package api

import (
	"time"

	"github.com/tjarratt/go-best-practices/domain"
)

//go:generate counterfeiter . PizzaDeliveryEstimator
type PizzaDeliveryEstimator interface {
	EstimatedDeliveryTime(domain.Pizza) time.Time
}
