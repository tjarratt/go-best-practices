package usecases

import (
	"time"

	"github.com/tjarratt/go-best-practices/api"
	"github.com/tjarratt/go-best-practices/domain"
)

//go:generate counterfeiter . OrderPizzaUseCase
type OrderPizzaUseCase interface {
	Execute(OrderPizzaRequest) (PizzaResponse, error)
}

type OrderPizzaRequest struct {
	Whom     string
	Address  string
	Dough    domain.DoughType
	Toppings []domain.Ingredient
}

type PizzaResponse struct {
	OrderNumber       int64
	EstimatedDelivery time.Duration
}

func NewOrderPizzaUseCase(
	pizzaRepository api.PizzaRepository,
	deliveryEstimator api.PizzaDeliveryEstimator,
) OrderPizzaUseCase {
	return orderPizzaUseCase{
		pizzaRepository:   pizzaRepository,
		deliveryEstimator: deliveryEstimator,
	}
}

type orderPizzaUseCase struct {
	pizzaRepository   api.PizzaRepository
	deliveryEstimator api.PizzaDeliveryEstimator
}

func (usecase orderPizzaUseCase) Execute(request OrderPizzaRequest) (PizzaResponse, error) {
	if request.Whom == "" {
		return PizzaResponse{}, &InvalidNameError{}
	}
	if request.Address == "" {
		return PizzaResponse{}, &InvalidAddressError{}
	}

	pizza, orderNumber, err := usecase.pizzaRepository.MakePizza(
		request.Dough,
		request.Toppings,
	)
	if err != nil {
		return PizzaResponse{}, err
	}

	deliveryTime := usecase.deliveryEstimator.EstimatedDeliveryTime(pizza)

	return PizzaResponse{
		OrderNumber:       orderNumber,
		EstimatedDelivery: deliveryTime,
	}, nil
}
