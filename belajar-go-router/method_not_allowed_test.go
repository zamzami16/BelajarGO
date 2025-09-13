package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestMethodNotAllowedHanlder(t *testing.T) {
	router := httprouter.New()
	router.GET("/resource", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "GET resource")
	})

	router.POST("/resource", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "POST resource")
	})
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Custom 405 Method Not Allowed")
	})

	req := httptest.NewRequest("PUT", "/resource", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusMethodNotAllowed, res.Code)
	assert.Equal(t, "Custom 405 Method Not Allowed", res.Body.String())
}