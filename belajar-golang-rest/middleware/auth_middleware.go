package middleware

import (
	"belajar-go-rest/helper"
	"belajar-go-rest/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

const API_KEY = "RAHASIA"

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-KEY") != API_KEY {
		writer.Header().Set("Content-Type", "Application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   "UNAUTHORIZED",
		}
		helper.WriteResponseBody(writer, webResponse)
		return
	}
	middleware.Handler.ServeHTTP(writer, request)
}
