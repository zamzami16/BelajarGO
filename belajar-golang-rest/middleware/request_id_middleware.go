package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type contextKey string

const RequestIDKey contextKey = "requestId"

func RequestIDMiddleware(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		requestID := uuid.New().String()
		ctx := context.WithValue(r.Context(), RequestIDKey, requestID)
		r = r.WithContext(ctx)
		w.Header().Set("X-Request-ID", requestID)
		h(w, r, ps)
	}
}

func GetRequestID(ctx context.Context) string {
	if requestID, ok := ctx.Value(RequestIDKey).(string); ok {
		return requestID
	}
	return "unknown"
}
