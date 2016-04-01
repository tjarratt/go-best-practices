package httpserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tjarratt/go-best-practices/domain"
)

//go:generate counterfeiter . OrderPizzaParamReader
type OrderPizzaParamReader interface {
	ReadParamsFromRequest(*http.Request) (domain.DoughType, []domain.Ingredient, error)
}

func NewOrderPizzaParamReader() OrderPizzaParamReader {
	return orderPizzaParamReader{}
}

type orderPizzaParamReader struct{}

func (paramReader orderPizzaParamReader) ReadParamsFromRequest(
	request *http.Request,
) (domain.DoughType, []domain.Ingredient, error) {
	bodyStr, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return 0, nil, err
	}

	var body map[string]interface{}
	err = json.Unmarshal(bodyStr, &body)
	if err != nil {
		return 0, nil, err
	}

	var (
		dough    domain.DoughType
		toppings []domain.Ingredient
	)

	switch body["dough"] {
	case "thin":
		dough = domain.Thin
	case "regular":
		dough = domain.Regular
	case "wheat":
		dough = domain.Wheat
	case "deep":
		dough = domain.Deep
	case "", nil:
		dough = domain.Regular
	default:
		return 0, nil, fmt.Errorf("unknown dough type '%s'", body["dough"])
	}

	toppings, err = parseToppingsFromBody(body)
	if err != nil {
		return 0, nil, fmt.Errorf("Bad request. The toppings are malformed!")
	}

	return dough, toppings, nil
}

func parseToppingsFromBody(body map[string]interface{}) ([]domain.Ingredient, error) {
	toppings := []domain.Ingredient{}
	if body["toppings"] == nil {
		return toppings, nil
	}

	requestedToppings, ok := body["toppings"].(interface{})
	if !ok {
		return nil, fmt.Errorf("Bad request. The toppings are malformed!")
	}

	requestedToppingsSlice, ok := requestedToppings.([]interface{})
	if !ok {
		return nil, fmt.Errorf("Bad request. The toppings are malformed!")
	}

	for _, topping := range requestedToppingsSlice {
		switch topping {
		case "pepperoni":
			toppings = append(toppings, domain.Pepperoni{})
		default:
			return nil, fmt.Errorf("unknown topping '%s'", topping)
		}
	}

	return toppings, nil
}
