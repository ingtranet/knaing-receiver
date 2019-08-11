package receiver

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestNriyReceiver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "nriy-receiver test Suite")
}