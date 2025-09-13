package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources/*
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	directory, err := fs.Sub(resources, "resources")
	if err != nil {
		t.Fatal("Failed to get subdirectory")
	}
	router.ServeFiles("/files/*filepath", http.FS(directory))

	req := httptest.NewRequest("GET", "/files/hello.txt", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	response := res.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "hello HttpRouter", string(body))
}