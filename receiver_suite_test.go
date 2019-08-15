package receiver

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestKnaingReceiver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "knaing-receiver test Suite")
}