package receiver

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
	"os"
)

var _ = Describe("stan client", func() {
	It("should publish well", func() {
		if os.Getenv("DOCKER_COMPOSE_TEST") != "1" {
			Skip("no docker-compose test")
		}
		config := viper.New()
		config.AutomaticEnv()
		config.Set("stan_client_id", "stan_client_test")
		client, err := newStanClient(config)
		Expect(err).Should(BeNil())

		err = client.Publish("test.subject", []byte("this is a test payload"))
		Expect(err).Should(BeNil())
	})
})