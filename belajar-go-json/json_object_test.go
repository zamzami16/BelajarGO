package main

import (
	"encoding/json"
	"testing"
)

type Address struct {
	Street string
	City   string
	Zip    string
}

type Customer struct {
	Name    string
	Age     int
	Email   string
	Address []Address
	Hobbies []string 
}

func NewCustomer(name string, age int, email string, address Address) Customer {
	return Customer{
		Name:  name,
		Age:   age,
		Email: email,
		Address: []Address{address},
		Hobbies: []string{},
	}
}

func TestObjectJson(t *testing.T) {
	customer := NewCustomer("Alice", 30, "alice@example.com", Address{
		Street: "123 Main St",
		City:   "Anytown",
		Zip:    "12345",
	})
	val, _ := json.Marshal(customer)
	t.Logf("JSON output: %s", val)
}

func TestUnmarshalJson(t *testing.T) {
	jsonData := `{"name":"Alice","age":30,"email":"alice@example.com","password":"secret"}`
	var customer Customer
	err := json.Unmarshal([]byte(jsonData), &customer)
	if err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", err)
	}
	t.Logf("Go struct output: %+v", customer)
}

func TestUnMarshalWithNotExsistsField(t *testing.T) {
	jsonData := `{"name":"Alice","age":30}`
	var customer Customer
	err := json.Unmarshal([]byte(jsonData), &customer)
	if err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", err)
	}
	t.Logf("Go struct output: %+v", customer)
}

func TestUnmarshalJsonWithHobbies(t *testing.T) {
	jsonData := `{"name":"Alice","age":30,"email":"alice@example.com","hobbies":["reading","traveling"]}`
	var customer Customer
	err := json.Unmarshal([]byte(jsonData), &customer)
	if err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", err)
	}
	t.Logf("Go struct output: %+v", customer)
}

func TestUnmarshalJsonWithAddress(t *testing.T) {
	jsonData := `{"name":"Alice","age":30,"email":"alice@example.com","address":[{"street":"123 Main St","city":"Anytown","zip":"12345"},{"street":"456 Oak St","city":"Othertown","zip":"67890"}]}`
	var customer Customer
	err := json.Unmarshal([]byte(jsonData), &customer)
	if err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", err)
	}
	t.Logf("Go struct output: %+v", customer)
}