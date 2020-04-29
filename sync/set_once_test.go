package sync_test

import (
	. "github.com/kunashu/simple.mud/sync"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SetOnce", func() {
	var set_once SetOnce

	BeforeEach(func() {
		set_once = SetOnce{}
	})

	Context("When setting a value", func() {
		It("Should only allow you to set the value once", func() {
			expected := "First Value"

			Expect(set_once.Set(expected)).To(BeTrue())

			actual, found := set_once.Get()
			Expect(found).To(BeTrue())
			Expect(actual).To(Equal(expected))

			Expect(set_once.Set("Other Value")).To(BeFalse())

			actual, found = set_once.Get()
			Expect(found).To(BeTrue())
			Expect(actual).To(Equal(expected))
		})

		It("Should be able to work synchronously", func() {
			cycles := 10
			output := make(chan bool, cycles)

			for i := 0; i < cycles; i++ {
				go func(v int) {
					output <- set_once.Set(v)
				}(i)
			}

			actual_fail := []bool{}
			for i := 0; i < cycles; i++ {
				select {
				case v := <-output:
					if !v {
						actual_fail = append(actual_fail, false)
					}
				}
			}

			Expect(actual_fail).To(HaveLen(cycles - 1))
		})

		It("Should be able to set nil", func() {
			Expect(set_once.Set(nil)).To(BeTrue())

			actual, found := set_once.Get()
			Expect(found).To(BeTrue())
			Expect(actual).To(BeNil())
		})
	})

	Context("When trying to get a value", func() {
		It("Should return nil and false if no value has been set", func() {
			actual, found := set_once.Get()
			Expect(found).To(BeFalse())
			Expect(actual).To(BeNil())
		})
	})
})
