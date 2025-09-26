package exception

import (
	"belajar-go-rest/helper"
	"belajar-go-rest/logging"
	"belajar-go-rest/middleware"
	"belajar-go-rest/model/web"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)

var globalLogger *logrus.Entry

func init() {
	loggerProvider := logging.NewLoggerProvider()
	globalLogger = loggerProvider.GetLogger("ExceptionHandler")
}

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {
	requestID := middleware.GetRequestID(request.Context())
	logger := getLoggerFromError(err)
	if logger != globalLogger {
		logger = logger.WithField("error_type", getErrorType(err))
	} else {
		logger = globalLogger.WithFields(logrus.Fields{
			"request_id": requestID,
			"error_type": getErrorType(err),
		})
	}

	errorString := getErrorString(err)

	logger.WithFields(logrus.Fields{
		"method": request.Method,
		"url":    request.URL.String(),
		"error":  errorString,
	}).Error("Exception occurred in request")

	if contextualNotFoundError(writer, err) {
		return
	}

	if notFoundError(writer, err) {
		return
	}

	if validationError(writer, err) {
		return
	}

	internalServerError(writer, err)
}

func getLoggerFromError(err any) *logrus.Entry {
	if contextualErr, ok := err.(ContextualNotFoundError); ok {
		return contextualErr.Logger
	}
	if contextualErr, ok := err.(ContextualError); ok {
		return contextualErr.Logger
	}
	return globalLogger
}

func getErrorString(err any) string {
	if err == nil {
		return "unknown error"
	}

	switch e := err.(type) {
	case ContextualNotFoundError:
		return e.NotFoundError.Error
	case NotFoundError:
		return e.Error
	case error:
		return e.Error()
	default:
		return fmt.Sprintf("%v", err)
	}
}

func contextualNotFoundError(writer http.ResponseWriter, err any) bool {
	exc, ok := err.(ContextualNotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "Application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exc.NotFoundError.Error,
		}

		helper.WriteResponseBody(writer, webResponse)

		return true
	}
	return false
}

func getErrorType(err any) string {
	switch err.(type) {
	case ContextualNotFoundError:
		return "CONTEXTUAL_NOT_FOUND"
	case NotFoundError:
		return "NOT_FOUND"
	case validator.ValidationErrors:
		return "VALIDATION_ERROR"
	default:
		return "INTERNAL_SERVER_ERROR"
	}
}

func validationError(writer http.ResponseWriter, err any) bool {
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

func notFoundError(writer http.ResponseWriter, err any) bool {
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

func internalServerError(writer http.ResponseWriter, err any) {
	writer.Header().Set("Content-Type", "Application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   getErrorString(err),
	}

	helper.WriteResponseBody(writer, webResponse)
}
