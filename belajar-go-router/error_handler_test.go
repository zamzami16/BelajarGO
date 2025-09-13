package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestErrorHandler(t *testing.T) {
	router := httprouter.New()
	router.GET("/panic", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		panic("something went wrong")
	})
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err any) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error: %v", err)
	}
	req := httptest.NewRequest("GET", "/panic", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code)
	assert.Equal(t, "Internal Server Error: something went wrong", res.Body.String())
}