package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type LogMiddleware struct{
	Handler http.Handler	
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, response *http.Request) {
	fmt.Printf("Request received: %s %s", response.Method, response.URL)
	middleware.Handler.ServeHTTP(writer, response)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Hello, world!")
	})

	req := httptest.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "Hello, world!", res.Body.String())
}

func TestLogMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Hello, world!")
	})

	loggedRouter := &LogMiddleware{Handler: router}
	req := httptest.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	
	loggedRouter.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "Hello, world!", res.Body.String())
}