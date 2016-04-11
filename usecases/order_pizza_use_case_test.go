package usecases_test

import (
	"errors"
	"time"

	"github.com/tjarratt/go-best-practices/domain"
	"github.com/tjarratt/go-best-practices/usecases/usecasesfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tjarratt/go-best-practices/usecases"
)

var _ = Describe("OrderPizzaUseCase", func() {
	var (
		subject           OrderPizzaUseCase
		pizzaFactory      *usecasesfakes.FakePizzaFactory
		deliveryEstimator *usecasesfakes.FakePizzaDeliveryEstimator

		request  OrderPizzaRequest
		response PizzaResponse
		err      error
	)

	BeforeEach(func() {
		pizzaFactory = new(usecasesfakes.FakePizzaFactory)
		deliveryEstimator = new(usecasesfakes.FakePizzaDeliveryEstimator)

		subject = NewOrderPizzaUseCase(pizzaFactory, deliveryEstimator)
	})

	JustBeforeEach(func() {
		response, err = subject.Execute(request)
	})

	Context("when the pizza order request is reasonable", func() {
		var expectedDeliveryTime = time.Minute

		BeforeEach(func() {
			request = OrderPizzaRequest{
				Whom:     "bone-thugz",
				Address:  "555 Crossroads Ave",
				Dough:    domain.Deep,
				Toppings: []domain.Ingredient{domain.Pepperoni{}},
			}

			deliveryEstimator.EstimatedDeliveryTimeReturns(expectedDeliveryTime)
			pizzaFactory.MakePizzaReturns(domain.Pizza{
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
			doughType, ingredients := pizzaFactory.MakePizzaArgsForCall(0)
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
				pizzaFactory.MakePizzaReturns(domain.Pizza{}, 0, errors.New("whoops!"))
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
			Expect(pizzaFactory.MakePizzaCallCount()).To(Equal(0))
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
			Expect(pizzaFactory.MakePizzaCallCount()).To(Equal(0))
		})
	})
})
