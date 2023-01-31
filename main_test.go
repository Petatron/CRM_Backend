package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unit Test Suite")
}

var _ = Describe("Main Test", func() {

	BeforeEach(func() {

	})
	AfterEach(func() {

	})

	Context("Customer Test", func() {
		It("Should create customer", func() {
			customer := createCustomer("1", "Andy", "Developer", "test", "404", false)
			Expect(customer.ID).To(Equal("1"))
			Expect(customer.Name).To(Equal("Andy"))
			Expect(customer.Role).To(Equal("Developer"))
			Expect(customer.Email).To(Equal("test"))
			Expect(customer.Phone).To(Equal("404"))
			Expect(customer.Contacted).To(Equal(false))
		})

		It("Should modify customer", func() {
			customer := createCustomer("1", "Andy", "Developer", "test", "404", false)
			customer.modifyCustomer("2", "Peter", "FBI", "new", "404=1", false)
			Expect(customer.ID).To(Equal("2"))
			Expect(customer.Name).To(Equal("Peter"))
			Expect(customer.Role).To(Equal("FBI"))
			Expect(customer.Email).To(Equal("new"))
			Expect(customer.Phone).To(Equal("404=1"))
			Expect(customer.Contacted).To(Equal(false))
		})

	})
})
