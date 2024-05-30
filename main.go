package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Customer struct {
	ID        string `json:"id,omitempty" bson:"id,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	Role      string `json:"role,omitempty" bson:"role,omitempty"`
	Email     string `json:"email,omitempty" bson:"email,omitempty"`
	Phone     int    `json:"phone,omitempty" bson:"phone,omitempty"`
	Contacted bool   `json:"contacted,omitempty" bson:"contacted,omitempty"`
}

var client *mongo.Client
var customerCollection *mongo.Collection

func initMongoDB() {
	var err error
	uri := "mongodb://localhost:27017"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	customerCollection = client.Database("CRM").Collection("customers")
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var customers []Customer
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := customerCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var customer Customer
		cursor.Decode(&customer)
		customers = append(customers, customer)
	}
	if len(customers) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "No customers available"}`))
		return
	}
	json.NewEncoder(w).Encode(customers)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := customerCollection.DeleteOne(ctx, bson.M{"id": params["id"]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if res.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Customer does not exist"}`))
		return
	}
	json.NewEncoder(w).Encode(res)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	var customer Customer
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := customerCollection.FindOne(ctx, bson.M{"id": params["id"]}).Decode(&customer)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Customer does not exist"}`))
		return
	}
	json.NewEncoder(w).Encode(customer)
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var customer Customer
	json.NewDecoder(r.Body).Decode(&customer)
	customer.ID = strconv.Itoa(int(time.Now().UnixNano())) // Generate a unique ID
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := customerCollection.InsertOne(ctx, customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	var customer Customer
	json.NewDecoder(r.Body).Decode(&customer)
	customer.ID = params["id"]
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := customerCollection.UpdateOne(
		ctx,
		bson.M{"id": params["id"]},
		bson.D{
			{"$set", customer},
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if res.MatchedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Customer does not exist"}`))
		return
	}
	json.NewEncoder(w).Encode(customer)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Main Page</h1>")
	fmt.Fprintf(w, "http://localhost:3000/customers To see all available customers")
	fmt.Fprintf(w, "<br>")
	fmt.Fprintf(w, "http://localhost:3000/customers/{id} To choose a specific customer")
}

func main() {
	initMongoDB()

	router := mux.NewRouter()
	router.HandleFunc("/", homepage)
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":3000", router))
}
