package usecases_test

import (
	"errors"
	"time"

	"github.com/tjarratt/go-best-practices/api/apifakes"
	"github.com/tjarratt/go-best-practices/domain"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tjarratt/go-best-practices/usecases"
)

var _ = Describe("OrderPizzaUseCase", func() {
	var (
		subject           OrderPizzaUseCase
		pizzaRepository   *apifakes.FakePizzaRepository
		deliveryEstimator *apifakes.FakePizzaDeliveryEstimator

		request  OrderPizzaRequest
		response PizzaResponse
		err      error
	)

	BeforeEach(func() {
		pizzaRepository = new(apifakes.FakePizzaRepository)
		deliveryEstimator = new(apifakes.FakePizzaDeliveryEstimator)

		subject = NewOrderPizzaUseCase(pizzaRepository, deliveryEstimator)
	})

	JustBeforeEach(func() {
		response, err = subject.Execute(request)
	})

	Context("when the pizza order request is reasonable", func() {
		var expectedDeliveryTime = time.Now()

		BeforeEach(func() {
			request = OrderPizzaRequest{
				Whom:     "bone-thugz",
				Address:  "555 Crossroads Ave",
				Dough:    domain.Deep,
				Toppings: []domain.Ingredient{domain.Pepperoni{}},
			}

			deliveryEstimator.EstimatedDeliveryTimeReturns(expectedDeliveryTime)
			pizzaRepository.MakePizzaReturns(domain.Pizza{
				Dough:    domain.Deep,
				Toppings: []domain.Ingredient{domain.Pepperoni{}},
			}, 666, nil)
		})

		It("should return a successful response when given a reasonable order request", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(response.OrderNumber).To(BeEquivalentTo(666))
			Expect(response.EstimatedDelivery).To(Equal(expectedDeliveryTime))
		})

		It("should make the pizza the customer ordered", func() {
			doughType, ingredients := pizzaRepository.MakePizzaArgsForCall(0)
			Expect(doughType).To(Equal(domain.Deep))
			Expect(ingredients).To(Equal([]domain.Ingredient{domain.Pepperoni{}}))
		})

		It("should ask its delivery estimator to estimate the time to deliver this pizza", func() {
			Expect(deliveryEstimator.EstimatedDeliveryTimeCallCount()).To(Equal(1))

			pizza := deliveryEstimator.EstimatedDeliveryTimeArgsForCall(0)
			Expect(pizza).To(Equal(domain.Pizza{
				Dough:    domain.Deep,
				Toppings: []domain.Ingredient{domain.Pepperoni{}},
			}))
		})

		Context("when submitting the order fails", func() {
			BeforeEach(func() {
				pizzaRepository.MakePizzaReturns(domain.Pizza{}, 0, errors.New("whoops!"))
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Context("when no name is provided", func() {
		BeforeEach(func() {
			request = OrderPizzaRequest{
				Whom:     "",
				Address:  "whoops",
				Dough:    domain.Deep,
				Toppings: []domain.Ingredient{},
			}
		})

		It("should return an error", func() {
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(&InvalidNameError{}))
		})

		It("should not try to make the pizza", func() {
			Expect(pizzaRepository.MakePizzaCallCount()).To(Equal(0))
		})
	})

	Context("when no name is provided", func() {
		BeforeEach(func() {
			request = OrderPizzaRequest{
				Whom:     "im-on-the-road-to-nowhere",
				Address:  "",
				Dough:    domain.Deep,
				Toppings: []domain.Ingredient{},
			}
		})

		It("should return an error", func() {
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(&InvalidAddressError{}))
		})

		It("should not try to make the pizza", func() {
			Expect(pizzaRepository.MakePizzaCallCount()).To(Equal(0))
		})
	})
})
