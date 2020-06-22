package fungo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFungo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fungo Suite")
}
