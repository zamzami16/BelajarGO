package main

import (
	"encoding/json"
	"testing"
)

func TestMapJson(t *testing.T) {
	customer := map[string]any{
		"name":  "Alice",
		"age":   30,
		"email": "alice@example.com",
		"address": []map[string]string{
			{
				"street": "123 Main St",
				"city":   "Anytown",
				"zip":    "12345",
			},
		},
	}

	val, _ := json.Marshal(customer)
	t.Logf("JSON output: %s", val)
}

func TestUnmarshalMapJson(t *testing.T) {
	jsonData := `{"name":"Alice","age":30,"email":"alice@example.com","address":[{"street":"123 Main St","city":"Anytown","zip":"12345"}]}`
	var customer map[string]any
	err := json.Unmarshal([]byte(jsonData), &customer)
	if err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", err)
	}
	t.Logf("Go struct output: %+v", customer)
}
