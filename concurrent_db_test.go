package fungo_test

import (
	"math/rand"
	"sync"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/williammartin/fungo"
)

var _ = Describe("InMemoryDB Concurrent Safe", func() {

	var inMemory *fungo.ConcurrentInMemoryDB

	BeforeEach(func() {
		inMemory = fungo.NewConcurrentInMemoryDB()
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

	It("serialises concurrent operations", func() {
		wg := &sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			go func() {
				wg.Add(1)
				val := String(10)
				inMemory.Set("foo", val)

				value, set := inMemory.Get("foo")
				Expect(set).To(BeTrue())
				Expect(value).To(Equal(val))
				wg.Done()
			}()
		}

		wg.Wait()
	})

})

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}
