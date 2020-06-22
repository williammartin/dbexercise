package dbexercise_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/williammartin/dbexercise"
)

var _ = Describe("InMemoryDB", func() {

	var inMemory *dbexercise.InMemoryDB

	BeforeEach(func() {
		inMemory = dbexercise.NewInMemoryDB()
	})

	It("stores key value mappings", func() {
		inMemory.Set("foo", "bar")

		value, set := inMemory.Get("foo")
		Expect(set).To(BeTrue())
		Expect(value).To(Equal("bar"))
	})

	It("returns false when a key has not been set", func() {
		_, set := inMemory.Get("foo")

		Expect(set).To(BeFalse())
	})

	When("transacting", func() {

		It("rolls back all transactions", func() {
			inMemory.BeginTransaction()

			inMemory.Set("foo", "bar")

			value, set := inMemory.Get("foo")
			Expect(set).To(BeTrue())
			Expect(value).To(Equal("bar"))

			inMemory.Rollback()

			_, set = inMemory.Get("foo")

			Expect(set).To(BeFalse())
		})

		It("commits all transactions", func() {
			inMemory.BeginTransaction()

			inMemory.Set("foo", "bar")

			value, set := inMemory.Get("foo")
			Expect(set).To(BeTrue())
			Expect(value).To(Equal("bar"))

			inMemory.Commit()

			value, set = inMemory.Get("foo")
			Expect(set).To(BeTrue())
			Expect(value).To(Equal("bar"))
		})

		It("errors if there is no transaction in progress when rolling back", func() {
			Expect(inMemory.Rollback()).To(MatchError("cannot rollback without transaction"))
		})

	})

})
