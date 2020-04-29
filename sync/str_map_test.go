package sync_test

import (
	"sync"

	. "github.com/kunashu/simple.mud/sync"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AsyncStrMap", func() {
	var test_map AsyncStrMap

	BeforeEach(func() {
		test_map = NewAsyncStrMap()
	})

	It("Should be able to set new entries to the map.", func() {
		expected := map[string]string{
			"A": "val-1", "B": "val-2", "C": "val-3",
		}

		for key, value := range expected {
			Expect(test_map.Set(key, value)).To(BeTrue())
		}

		Expect(test_map.Len()).To(Equal(len(expected)))

		for key, value := range expected {
			actual, ok := test_map.Get(key)

			Expect(ok).To(BeTrue())
			Expect(actual).To(Equal(value))
		}
	})

	It("Should not insert new entries with the same key.", func() {
		key := "test"
		expected := "test-val"

		Expect(test_map.Set(key, expected)).To(BeTrue())
		Expect(test_map.Set(key, "other")).To(BeFalse())

		actual, ok := test_map.Get(key)
		Expect(ok).To(BeTrue())
		Expect(actual).To(Equal(expected))
	})

	It("Should return nil and false when the key isn't found", func() {
		actual, ok := test_map.Get("invalid")
		Expect(ok).To(BeFalse())
		Expect(actual).To(BeNil())
	})

	It("Should be able to work synchronously", func() {
		expected := map[string]string{
			"A": "val-1", "B": "val-2", "C": "val-3",
			"D": "val-4", "E": "val-5", "F": "val-6",
			"G": "val-4", "H": "val-5", "I": "val-6",
			"J": "val-4", "K": "val-5", "L": "val-6",
			"M": "val-4", "N": "val-5", "O": "val-6",
			"P": "val-4", "Q": "val-5", "R": "val-6",
		}
		start, done := sync.WaitGroup{}, sync.WaitGroup{}
		start.Add(1)

		for key, value := range expected {
			done.Add(1)
			go func(key, value string) {
				defer done.Done()

				start.Wait()

				Expect(test_map.Set(key, value)).To(BeTrue())

				actual, ok := test_map.Get(key)
				Expect(ok).To(BeTrue())
				Expect(actual).To(Equal(value))
			}(key, value)
		}

		start.Done()
		done.Wait()

		Expect(test_map.Len()).To(Equal(len(expected)))
	})
})
