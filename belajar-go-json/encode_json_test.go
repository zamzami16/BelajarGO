package main

import (
	"encoding/json"
	"log"
	"testing"
)

func logJson(data any) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return
	}
	log.Printf("JSON output: %s", jsonData)
}

type CustomStruct struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func New(field1 string, field2 int) CustomStruct {
	return CustomStruct{Field1: field1, Field2: field2}
}

func TestLogJson(t *testing.T) {
	logJson(1)
	logJson("string")
	logJson(true)
	logJson([]string{"item1", "item2"})
	logJson(nil)
	logJson(New("example", 42))
	logJson(map[string]string{"key": "value"})
}