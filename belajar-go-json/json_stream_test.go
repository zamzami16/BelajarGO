package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestDecodeJsonStream(testing *testing.T) {
	reader, _ := os.Open("customer.json")
	defer reader.Close()

	decoder := json.NewDecoder(reader)
	var customers []Customer
	err := decoder.Decode(&customers)
	if err != nil && err != io.EOF {
		testing.Fatalf("Error decoding JSON: %v", err)
	}

	testing.Logf("Decoded customers: %+v", customers)
}

func TestEncodeJsonStream(testing *testing.T) {
	var customers []Customer
	// Populate the customers slice as needed
	for i := 0; i < 1_000; i++ {
		customers = append(customers, NewCustomer(
			"Customer"+fmt.Sprint(i+1),
			20+i,
			"customer"+fmt.Sprint(i+1)+"@example.com",
			Address{
				Street: "123 Main St",
				City:   "Anytown",
				Zip:    "12345",
			},
		))
	}

	writer, _ := os.Create("output_customers.json")
	defer writer.Close()

	encoder := json.NewEncoder(writer)
	err := encoder.Encode(customers)
	if err != nil {
		testing.Fatalf("Error encoding JSON: %v", err)
	}
}
