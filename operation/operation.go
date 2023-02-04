package operation

import (
	cs "CRM_backend/customer" // import the customer package
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// Customers Mock database of customers
var Customers = make(map[string]cs.Customer)

// GetCustomers returns all customers
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(Customers)
	if err != nil {
		panic(err)
	}
}

// GetCustomer returns a single customer
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if _, exist := Customers[id]; exist {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(Customers[id])
		if err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(Customers)
		if err != nil {
			panic(err)
		}
	}
}

// AddCustomer adds a new customer
func AddCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newEntry *cs.Customer

	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &newEntry)
	if err != nil {
		panic(err)
	}

	if _, exist := Customers[newEntry.ID]; exist {
		w.WriteHeader(http.StatusConflict)
	} else {
		Customers[newEntry.ID] = *newEntry
		w.WriteHeader(http.StatusCreated)
	}

	err = json.NewEncoder(w).Encode(Customers)
	if err != nil {
		panic(err)
	}
}

// UpdateCustomer updates a customer info
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	newEntry := new(cs.Customer)

	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &newEntry)
	if err != nil {
		panic(err)
	}

	if _, exist := Customers[newEntry.ID]; exist {
		newEntry.ModifyCustomer(
			newEntry.ID,
			newEntry.Name,
			newEntry.Role,
			newEntry.Email,
			newEntry.Phone,
			newEntry.Contacted)
		Customers[newEntry.ID] = *newEntry
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(Customers)
		if err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(Customers)
		if err != nil {
			panic(err)
		}
	}

}

// DeleteCustomer deletes a customer
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if _, exist := Customers[id]; exist {
		delete(Customers, id)
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(Customers)
		if err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(Customers)
		if err != nil {
			panic(err)
		}
	}
}
