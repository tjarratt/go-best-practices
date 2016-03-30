package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/pizza", OrderPizzaHandler).Methods("POST")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err.Error())
	}
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	println("handling a home request")
	writer.Write([]byte("handling a home request"))
}

func OrderPizzaHandler(writer http.ResponseWriter, request *http.Request) {
	println("handling an orderPizza request")
	writer.Write([]byte("handling an orderPizza request"))
}
