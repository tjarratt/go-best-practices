package httpserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tjarratt/go-best-practices/usecases"
)

func NewOrderPizzaHandler(
	useCase usecases.OrderPizzaUseCase,
	paramReader OrderPizzaParamReader,
) http.Handler {
	return orderPizzaHandler{
		useCase:     useCase,
		paramReader: paramReader,
	}
}

type orderPizzaHandler struct {
	useCase     usecases.OrderPizzaUseCase
	paramReader OrderPizzaParamReader
}

func (handler orderPizzaHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	bodyStr, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var body map[string]interface{}
	err = json.Unmarshal(bodyStr, &body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	doughType, toppings, err := handler.paramReader.ReadParamsFromRequest(request)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := handler.useCase.Execute(usecases.OrderPizzaRequest{
		Whom:     body["name"].(string),
		Address:  body["address"].(string),
		Dough:    doughType,
		Toppings: toppings,
	})

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	estimatedMinutes := response.EstimatedDelivery.Minutes()
	message := fmt.Sprintf(successMessage, int(estimatedMinutes))
	writer.Write([]byte(message))
}

const successMessage string = "Your pizza will be ready in %d minutes"
