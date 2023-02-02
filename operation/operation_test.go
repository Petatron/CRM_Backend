package operation

import (
	cs "CRM_backend/customer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unit Test for Operation")
}

var _ = Describe("Operation", func() {

	BeforeEach(func() {

	})
	AfterEach(func() {

	})

	Context("customer functions", func() {
		It("Test GetCustomers", func() {
			ca := cs.CreateCustomer("1", "Andy", "Developer", "S", "S", true)
			cb := cs.CreateCustomer("2", "Peter", "Developer", "S", "S", true)
			Customers[ca.ID] = *ca
			Customers[cb.ID] = *cb

		})

	})
})
