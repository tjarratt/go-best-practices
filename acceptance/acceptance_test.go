package acceptance_test

import (
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var pathToServerExecutable string

var _ = BeforeSuite(func() {
	var err error
	pathToServerExecutable, err = gexec.Build("github.com/tjarratt/go-best-practices")
	Expect(err).ToNot(HaveOccurred())
})

var _ = Describe("The Pizza HTTP Server", func() {
	var session *gexec.Session

	BeforeEach(func() {
		command := exec.Command(pathToServerExecutable)

		var err error
		session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		session.Kill()
	})

	It("should handle http requests on port 8080", func() {
		Eventually(session).Should(gbytes.Say("listening on port 8080"))

		request := strings.NewReader(`{
						"name":     "Obama",
						"address":  "1600 Pennsylvania Ave.",
						"dough":    "thin",
						"toppings": []
		}`)

		response, err := http.Post("http://localhost:8080/pizza", "encoding/json", request)
		Expect(err).ToNot(HaveOccurred())

		body, err := ioutil.ReadAll(response.Body)
		Expect(err).ToNot(HaveOccurred())
		Expect(string(body)).To(Equal("Your pizza will be ready in 30 minutes"))
		Expect(response.StatusCode).To(Equal(http.StatusOK))
	})
})
