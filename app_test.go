package receiver

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("app", func() {
	It("should be initialized well", func() {
		if os.Getenv("DOCKER_COMPOSE_TEST") != "1" {
			Skip("no docker-compose test")
		}
		app := NewApp()
		Expect(app.Err()).Should(BeNil())
	})
})