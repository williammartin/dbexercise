package dbexercise_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDbexercise(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dbexercise Suite")
}
