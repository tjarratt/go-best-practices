package acceptance_test

import (
	"os/exec"

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
	It("should start listening on port 8080 when started", func() {
		command := exec.Command(pathToServerExecutable)
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())

		Eventually(session).Should(gbytes.Say("listening on port 8080"))
	})
})
