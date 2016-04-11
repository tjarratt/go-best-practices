package usecases

import (
	"time"

	"github.com/tjarratt/go-best-practices/domain"
)

//go:generate counterfeiter . OrderPizzaUseCase
type OrderPizzaUseCase interface {
	Execute(OrderPizzaRequest) (PizzaResponse, error)
}

//go:generate counterfeiter . PizzaFactory
type PizzaFactory interface {
	MakePizza(domain.DoughType, []domain.Ingredient) (domain.Pizza, int64, error)
}

//go:generate counterfeiter . PizzaDeliveryEstimator
type PizzaDeliveryEstimator interface {
	EstimatedDeliveryTime(domain.Pizza) time.Duration
}

func NewOrderPizzaUseCase(
	pizzaFactory PizzaFactory,
	deliveryEstimator PizzaDeliveryEstimator,
) OrderPizzaUseCase {
	return orderPizzaUseCase{
		pizzaFactory:      pizzaFactory,
		deliveryEstimator: deliveryEstimator,
	}
}

type orderPizzaUseCase struct {
	pizzaFactory      PizzaFactory
	deliveryEstimator PizzaDeliveryEstimator
}

func (usecase orderPizzaUseCase) Execute(request OrderPizzaRequest) (PizzaResponse, error) {
	if request.Whom == "" {
		return PizzaResponse{}, &InvalidNameError{}
	}
	if request.Address == "" {
		return PizzaResponse{}, &InvalidAddressError{}
	}

	pizza, orderNumber, err := usecase.pizzaFactory.MakePizza(
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
