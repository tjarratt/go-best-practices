package httpserver_test

import (
	"net/http"
	"strings"

	"github.com/tjarratt/go-best-practices/domain"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tjarratt/go-best-practices/httpserver"
)

var _ = Describe("OrderPizzaParamReader", func() {
	var subject OrderPizzaParamReader
	var requestBody string

	var dough domain.DoughType
	var toppings []domain.Ingredient
	var err error

	BeforeEach(func() {
		subject = NewOrderPizzaParamReader()
	})

	subjectAction := func() {
		body := strings.NewReader(requestBody)
		request, _ := http.NewRequest("POST", "http://example.com/pizza", body)
		dough, toppings, err = subject.ReadParamsFromRequest(request)
	}

	Describe("dough", func() {
		It("is thin when the user requests thin dough", func() {
			requestBody = thinBody
			subjectAction()
			Expect(err).ToNot(HaveOccurred())

			Expect(dough).To(Equal(domain.Thin))
		})

		It("is regular by default", func() {
			requestBody = emptyBody
			subjectAction()
			Expect(err).ToNot(HaveOccurred())

			Expect(dough).To(Equal(domain.Regular))
		})

		It("can be requested to be wheat", func() {
			requestBody = wheatBody
			subjectAction()
			Expect(err).ToNot(HaveOccurred())

			Expect(dough).To(Equal(domain.Wheat))
		})

		It("can be deep if the user likes deep dish", func() {
			requestBody = deepBody
			subjectAction()
			Expect(err).ToNot(HaveOccurred())

			Expect(dough).To(Equal(domain.Deep))
		})
	})

	Describe("toppings", func() {
		It("knows how to parse pepperoni", func() {
			requestBody = pepperoniBody
			subjectAction()
			Expect(err).ToNot(HaveOccurred())

			Expect(toppings).To(ContainElement(domain.Pepperoni{}))
		})

		It("defaults to an empty list if no ingredients are requested", func() {
			requestBody = emptyBody
			subjectAction()
			Expect(err).ToNot(HaveOccurred())

			Expect(toppings).To(BeEmpty())
		})
	})

	It("returns an error if unknown dough is requested", func() {
		requestBody = garbageDough
		subjectAction()

		Expect(err).To(HaveOccurred())
	})

	It("returns an error if an unknown ingredient is requested", func() {
		requestBody = garbageDough
		subjectAction()

		Expect(err).To(HaveOccurred())
	})
})

const (
	emptyBody      string = `{}`
	thinBody       string = `{"dough": "thin"}`
	wheatBody      string = `{"dough": "wheat"}`
	deepBody       string = `{"dough": "deep"}`
	pepperoniBody  string = `{"toppings": ["pepperoni"]}`
	garbageDough   string = `{"dough": "garbage"}`
	garbageTopping string = `{"toppings": ["whooops"]}`
)
