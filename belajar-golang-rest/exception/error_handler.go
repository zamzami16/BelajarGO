package exception

import (
	"belajar-go-rest/helper"
	"belajar-go-rest/model/web"
	"net/http"

	"github.com/go-playground/validator"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	println("[DEBUG] ErrorHandler called with error:", err)
	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationError(writer http.ResponseWriter, _ *http.Request, err interface{}) bool {
	exc, ok := err.(validator.ValidationErrors)

	if ok {
		writer.Header().Set("Content-Type", "Application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exc.Error(),
		}

		helper.WriteResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, _ *http.Request, err interface{}) bool {
	exc, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "Application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exc.Error,
		}

		helper.WriteResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, _ *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "Application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteResponseBody(writer, webResponse)
}
