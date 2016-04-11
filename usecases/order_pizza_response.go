package usecases

import "time"

type PizzaResponse struct {
	OrderNumber       int64
	EstimatedDelivery time.Duration
}
