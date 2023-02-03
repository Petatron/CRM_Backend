package main

import (
	cs "CRM_backend/customer"
	op "CRM_backend/operation"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     int32 `json:"phone"`
	Contacted bool   `json:"contacted"`
}

// createCustomer creates a new customer
func createCustomer(ID, Name, Role, Email string, Phone int32, Contacted bool) (result *Customer) {
	customer := Customer{
		ID:        ID,
		Name:      Name,
		Role:      Role,
		Email:     Email,
		Phone:     Phone,
		Contacted: Contacted,
	}
	result = &customer
	return result
}

// modifyCustomer modifies an existing customer
func (c *Customer) modifyCustomer(ID, Name, Role, Email, Phone string, Contacted bool) {
	c.ID = ID
	c.Name = Name
	c.Role = Role
	c.Email = Email
	c.Phone = Phone
	c.Contacted = Contacted
}

var customers = map[string]string{
	"1": "Andy",
	"2": "Peter",
	"3": "Gabriella",
	"4": "Jordy",
}

func main() {
	ca := cs.CreateCustomer("1", "Andy", "Developer", "S", 320, true)
	cb := cs.CreateCustomer("2", "Peter", "Developer", "S", 408, true)
	op.Customers[ca.ID] = *ca
	op.Customers[cb.ID] = *cb

	fileServer := http.FileServer(http.Dir("./static"))

	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/", fileServer)

	fmt.Println("Server is starting on port 3000. You can access it on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
