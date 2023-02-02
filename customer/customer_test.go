package customer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unit Test for Customer")
}

var _ = Describe("Customer Test", func() {

	BeforeEach(func() {

	})
	AfterEach(func() {

	})

	Context("Customer function test", func() {
		It("Test CreateCustomer", func() {
			ca := CreateCustomer("1", "Andy", "Developer", "S", "S", true)

			Expect(ca.ID).To(Equal("1"))
			Expect(ca.Name).To(Equal("Andy"))
			Expect(ca.Role).To(Equal("Developer"))
			Expect(ca.Email).To(Equal("S"))
			Expect(ca.Phone).To(Equal("S"))
			Expect(ca.Contacted).To(Equal(true))
		})

		It("Test ModifyCustomer", func() {
			ca := CreateCustomer("1", "Andy", "Developer", "S", "S", true)
			ca.ModifyCustomer("2", "Peter", "FBI", "S", "S", true)

			Expect(ca.ID).To(Equal("2"))
			Expect(ca.Name).To(Equal("Peter"))
			Expect(ca.Role).To(Equal("FBI"))
			Expect(ca.Email).To(Equal("S"))
			Expect(ca.Phone).To(Equal("S"))
			Expect(ca.Contacted).To(Equal(true))
		})
	})
})
