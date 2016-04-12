package httpserver_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/tjarratt/go-best-practices/domain"
	"github.com/tjarratt/go-best-practices/httpserver/httpserverfakes"
	"github.com/tjarratt/go-best-practices/usecases"
	"github.com/tjarratt/go-best-practices/usecases/usecasesfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tjarratt/go-best-practices/httpserver"
)

var _ = Describe("OrderPizzaHandler", func() {
	var (
		subject     http.Handler
		useCase     *usecasesfakes.FakeOrderPizzaUseCase
		paramReader *httpserverfakes.FakeOrderPizzaParamReader

		responseWriter *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		responseWriter = httptest.NewRecorder()

		useCase = new(usecasesfakes.FakeOrderPizzaUseCase)
		paramReader = new(httpserverfakes.FakeOrderPizzaParamReader)
		subject = NewOrderPizzaHandler(useCase, paramReader)
	})

	Describe("a successful request", func() {
		BeforeEach(func() {
			useCase.ExecuteReturns(usecases.PizzaResponse{
				EstimatedDelivery: time.Minute * 666}, nil)
			paramReader.ReadParamsFromRequestReturns(
				OrderPizzaParams{
					Name:     "thuggish-ruggish-bone",
					Address:  "thug-mansion",
					Dough:    domain.Thin,
					Toppings: []domain.Ingredient{domain.Pepperoni{}},
				}, nil)
			body := strings.NewReader(`
{
	"name": "thuggish-ruggish-bone",
	"address": "thug-mansion",
	"dough": "thin",
	"toppings": ["pepperoni"]
}`)
			request, _ := http.NewRequest("POST", "http://example.com/pizza", body)
			subject.ServeHTTP(responseWriter, request)
		})

		It("should call its use case when handling a request", func() {
			Expect(useCase.ExecuteCallCount()).To(Equal(1))
			Expect(useCase.ExecuteArgsForCall(0)).To(Equal(usecases.OrderPizzaRequest{
				Whom:     "thuggish-ruggish-bone",
				Address:  "thug-mansion",
				Dough:    domain.Thin,
				Toppings: []domain.Ingredient{domain.Pepperoni{}},
			}))
		})

		It("should tell the user how long they should be waiting", func() {
			Expect(responseWriter.Body.String()).To(Equal("Your pizza will be ready in 666 minutes"))
		})
	})

	Context("when the user requests some unknown dough", func() {
		BeforeEach(func() {
			paramReader.ReadParamsFromRequestReturns(OrderPizzaParams{}, errors.New("whoops"))
			body := strings.NewReader(`
{
	"name": "thuggish-ruggish-bone",
	"address": "thug-mansion",
	"dough": "garbage",
	"toppings": ["pepperoni"]
}`)
			request, _ := http.NewRequest("POST", "http://example.com/pizza", body)
			subject.ServeHTTP(responseWriter, request)
		})

		It("should set the response status code to 4xx", func() {
			Expect(responseWriter.Code).To(Equal(http.StatusBadRequest))
		})

		It("should inform the user their request was bad", func() {
			Expect(responseWriter.Body.String()).ToNot(BeEmpty())
		})
	})

	Context("when the user requests some totally unknown ingredient", func() {
		BeforeEach(func() {
			paramReader.ReadParamsFromRequestReturns(OrderPizzaParams{}, errors.New("whoops"))
			body := strings.NewReader(`
{
	"name": "thuggish-ruggish-bone",
	"address": "thug-mansion",
	"dough": "thin",
	"toppings": ["garbage"]
}`)
			request, _ := http.NewRequest("POST", "http://example.com/pizza", body)
			subject.ServeHTTP(responseWriter, request)
		})

		It("should set the response status code to 4xx", func() {
			Expect(responseWriter.Code).To(Equal(http.StatusBadRequest))
		})

		It("should inform the user their request was bad", func() {
			Expect(responseWriter.Body.String()).ToNot(BeEmpty())
		})
	})

	Context("when executing the use case returns a validation error", func() {
		BeforeEach(func() {
			useCase.ExecuteReturns(usecases.PizzaResponse{}, errors.New("ruh roh"))
			body := strings.NewReader(`
{
	"name": "thuggish-ruggish-bone",
	"address": "thug-mansion",
	"dough": "thin",
	"toppings": []
}`)
			request, _ := http.NewRequest("POST", "http://example.com/pizza", body)
			subject.ServeHTTP(responseWriter, request)
		})

		It("should set the response status code to 4xx", func() {
			Expect(responseWriter.Code).To(Equal(http.StatusBadRequest))
		})

		It("should inform the user their request was bad", func() {
			Expect(responseWriter.Body.String()).ToNot(BeEmpty())
		})
	})
})
