package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestNotFoundHandler(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Custom 404 Not Found")
	})

	req := httptest.NewRequest("GET", "/nonexistent", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNotFound, res.Code)
	assert.Equal(t, "Custom 404 Not Found", res.Body.String())
}