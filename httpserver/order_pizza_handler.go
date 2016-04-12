package httpserver

import (
	"fmt"
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

	params, err := handler.paramReader.ReadParamsFromRequest(request)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
		return
	}

	response, err := handler.useCase.Execute(usecases.OrderPizzaRequest{
		Whom:     params.Name,
		Address:  params.Address,
		Dough:    params.Dough,
		Toppings: params.Toppings,
	})

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
		return
	}

	estimatedMinutes := response.EstimatedDelivery.Minutes()
	message := fmt.Sprintf(successMessage, int(estimatedMinutes))
	writer.Write([]byte(message))
}

const successMessage string = "Your pizza will be ready in %d minutes"
