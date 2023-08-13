package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var customers = struct {
	sync.RWMutex
	m map[string]Customer
}{
	m: map[string]Customer{
		"1": {"1", "Alice", "Engineer", "alice@example.com", 1234567890, false},
		"2": {"2", "Bob", "Manager", "bob@example.com", 1234567891, true},
		"3": {"3", "Charlie", "Director", "charlie@example.com", 1234567892, false},
	},
}

func main() {
	http.HandleFunc("/customers", customersHandler)
	http.HandleFunc("/customers/", customerHandler)
	http.ListenAndServe(":8080", nil)
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getCustomers(w, r)
	case "POST":
		addCustomer(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func customerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getCustomer(w, r)
	case "DELETE":
		deleteCustomer(w, r)
	case "PUT":
		updateCustomer(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	customers.RLock()
	defer customers.RUnlock()

	custs := make([]Customer, 0, len(customers.m))
	for _, value := range customers.m {
		custs = append(custs, value)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(custs)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	customers.RLock()
	defer customers.RUnlock()

	id := strings.TrimPrefix(r.URL.Path, "/customers/")
	customer, exists := customers.m[id]

	if !exists {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customers.Lock()
	customer.ID = fmt.Sprint(len(customers.m) + 1)
	customers.m[customer.ID] = customer
	customers.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/customers/")
	customers.Lock()
	defer customers.Unlock()

	if _, exists := customers.m[id]; !exists {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	var customer Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customer.ID = id
	customers.m[id] = customer

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	customers.Lock()
	defer customers.Unlock()

	id := strings.TrimPrefix(r.URL.Path, "/customers/")
	if _, exists := customers.m[id]; !exists {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	delete(customers.m, id)
	w.WriteHeader(http.StatusOK)
}