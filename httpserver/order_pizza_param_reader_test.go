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
	var (
		subject     OrderPizzaParamReader
		requestBody string

		params OrderPizzaParams
		err    error
	)

	BeforeEach(func() {
		subject = NewOrderPizzaParamReader()
	})

	subjectAction := func() {
		body := strings.NewReader(requestBody)
		request, _ := http.NewRequest("POST", "http://example.com/pizza", body)
		params, err = subject.ReadParamsFromRequest(request)
	}

	Describe("dough", func() {
		It("is thin when the user requests thin dough", func() {
			requestBody = thinBody
			subjectAction()

			Expect(err).ToNot(HaveOccurred())
			Expect(params.Dough).To(Equal(domain.Thin))
		})

		It("is regular by default", func() {
			requestBody = emptyBody
			subjectAction()

			Expect(err).ToNot(HaveOccurred())
			Expect(params.Dough).To(Equal(domain.Regular))
		})

		It("can be requested to be wheat", func() {
			requestBody = wheatBody
			subjectAction()

			Expect(err).ToNot(HaveOccurred())
			Expect(params.Dough).To(Equal(domain.Wheat))
		})

		It("can be deep if the user likes deep dish", func() {
			requestBody = deepBody
			subjectAction()

			Expect(err).ToNot(HaveOccurred())
			Expect(params.Dough).To(Equal(domain.Deep))
		})
	})

	Describe("toppings", func() {
		It("knows how to parse pepperoni", func() {
			requestBody = pepperoniBody
			subjectAction()

			Expect(err).ToNot(HaveOccurred())
			Expect(params.Toppings).To(ContainElement(domain.Pepperoni{}))
		})

		It("defaults to an empty list if no ingredients are requested", func() {
			requestBody = emptyBody
			subjectAction()

			Expect(err).ToNot(HaveOccurred())
			Expect(params.Toppings).To(BeEmpty())
		})
	})

	Describe("the name and address", func() {
		It("is available when present", func() {
			requestBody = completeBody
			subjectAction()

			Expect(err).ToNot(HaveOccurred())
			Expect(params.Name).To(Equal("David Byrne"))
			Expect(params.Address).To(Equal("road-to-nowhere"))
		})

		It("returns an error if the name is not a string", func() {
			requestBody = badName
			subjectAction()

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Bad request. Name is malformed"))
		})

		It("returns an error if the address is not a string", func() {
			requestBody = badAddress
			subjectAction()

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Bad request. Address is malformed"))
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

	Describe("when the request body is not JSON", func() {
		It("should return an error informing the client of the failing validation", func() {
			requestBody = "absolute-garbage"
			subjectAction()

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Expected the request body to be JSON but it was 'absolute-garbage'"))
		})
	})
})

const (
	emptyBody      string = `{"name": "", "address": ""}`
	thinBody       string = `{"name": "", "address": "", "dough": "thin"}`
	wheatBody      string = `{"name": "", "address": "", "dough": "wheat"}`
	deepBody       string = `{"name": "", "address": "", "dough": "deep"}`
	pepperoniBody  string = `{"name": "", "address": "", "toppings": ["pepperoni"]}`
	garbageDough   string = `{"name": "", "address": "", "dough": "garbage"}`
	garbageTopping string = `{"name": "", "address": "", "toppings": ["whooops"]}`
	completeBody   string = `{"name": "David Byrne", "address": "road-to-nowhere", "dough": "thin"}`
	badName        string = `{"name": 5, "address": "", "dough": "thin"}`
	badAddress     string = `{"name": "", "address": 5, "dough": "thin"}`
)
