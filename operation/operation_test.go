package operation

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
			req, err := http.NewRequest("GET", "/customers", nil)
			if err != nil {
				panic(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetCustomers)
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				_ = fmt.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			// Check the response body is what we expect.
			expected := `ewq`
			if rr.Body.String() != expected {
				_ = fmt.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}

		})

		It("Test GetCustomer", func() {

		})

	})
})
