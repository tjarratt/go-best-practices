package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/tjarratt/go-best-practices/api"
	"github.com/tjarratt/go-best-practices/httpserver"
	"github.com/tjarratt/go-best-practices/usecases"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)

	pizzaRepository := api.NewPizzaRepository()
	deliveryEstimator := api.NewPizzaDeliveryEstimator()
	orderPizzaHandler := httpserver.NewOrderPizzaHandler(
		usecases.NewOrderPizzaUseCase(
			pizzaRepository,
			deliveryEstimator,
		),
		httpserver.NewOrderPizzaParamReader(),
	)
	router.Handle("/pizza", orderPizzaHandler).Methods("POST")

	fmt.Fprintln(os.Stdout, "listening on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err.Error())
	}
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	println("handling a home request")
	writer.Write([]byte("handling a home request"))
}
