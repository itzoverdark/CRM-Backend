package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID        string
	Name      string
	Role      string
	Email     string
	Phone     int
	Contacted bool
}

var customers = []Customer{
	{ID: "1", Name: "James", Role: "Teacher", Email: "james@gmail.com", Phone: 2025550988, Contacted: true},
	{ID: "2", Name: "John", Role: "Lawyer", Email: "John@gmail.com", Phone: 2025550533, Contacted: false},
	{ID: "3", Name: "Tom", Role: "Software Developer", Email: "Tom@gmail.com", Phone: 2025550387, Contacted: true},
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	if len(customers) > 0 {
		json.NewEncoder(w).Encode(customers)
		return
	}
	w.WriteHeader(404)
	w.Write([]byte(`{"message": "No customers available"}`))
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for index, customer := range customers {
		if len(params["id"]) > 0 && customer.ID == params["id"] {
			customers = append(customers[:index], customers[index+1:]...)
			json.NewEncoder(w).Encode(customer)
			return
		}
	}

	w.WriteHeader(404)
	w.Write([]byte(`{"message": "Customer does not exist"}`))

}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for _, customer := range customers {
		if len(params["id"]) > 0 && customer.ID == params["id"] {
			json.NewEncoder(w).Encode(customer)
			return
		}
	}

	w.WriteHeader(404)
	w.Write([]byte(`{"message": "Customer does not exist"}`))

}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var customer Customer
	json.NewDecoder(r.Body).Decode(&customer)
	customer.ID = strconv.Itoa(len(customers) + 1)
	w.WriteHeader(201)
	customers = append(customers, customer)
	json.NewEncoder(w).Encode(customer)

}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for index, customer := range customers {
		if customer.ID == params["id"] {
			customers = append(customers[:index], customers[index+1:]...)
			var customer Customer
			json.NewDecoder(r.Body).Decode(&customer)
			customer.ID = params["id"]
			customers = append(customers, customer)
			json.NewEncoder(w).Encode(customer)
			return
		}
	}

	w.WriteHeader(404)
	w.Write([]byte(`{"message": "Customer does not exist"}`))

}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Main Page</h1>")
	fmt.Fprintf(w, "http://localhost:3000/customers To see all available customers")
	fmt.Fprintf(w, "<br>")
	fmt.Fprintf(w, "http://localhost:3000/customers/{id} To choose a specific customer")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homepage)
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")

	fmt.Println("Server started")
	http.ListenAndServe(":3000", router)
}
