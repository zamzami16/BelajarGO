package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestParam(t *testing.T) {
	router := httprouter.New()
	router.GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		name := ps.ByName("name")
		fmt.Fprintf(w, "Hello, %s!", name)
	})

	req, _ := http.NewRequest("GET", "/hello/John", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "Hello, John!", res.Body.String())
}

func TestParamWithMultiParam(t *testing.T) {
	router := httprouter.New()
	router.GET("/user/:user_id/contact/:contact_id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		userId := ps.ByName("user_id")
		contactId := ps.ByName("contact_id")
		fmt.Fprintf(w, "User ID: %s, Contact ID: %s", userId, contactId)
	})

	req, _ := http.NewRequest("GET", "/user/42/contact/100", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "User ID: 42, Contact ID: 100", res.Body.String())
}

// catch all should be the last route and with * prefix
func TestParamCatchAll(t *testing.T) {
	router := httprouter.New()
	router.GET("/files/*filepath", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		filepath := ps.ByName("filepath")
		fmt.Fprintf(w, "Filepath: %s", filepath)
	})

	req := httptest.NewRequest("GET", "/files/images/photo.jpg", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "Filepath: /images/photo.jpg", res.Body.String())
}