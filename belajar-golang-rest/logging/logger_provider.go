package logging

import (
	"belajar-go-rest/middleware"
	"context"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	instance *logrus.Logger
	once     sync.Once
)

// LoggerProvider is a singleton provider for logrus logger
type LoggerProvider struct {
	logger *logrus.Logger
}

// NewLoggerProvider creates a singleton logger provider
func NewLoggerProvider() *LoggerProvider {
	once.Do(func() {
		instance = logrus.New()
		instance.SetFormatter(&logrus.JSONFormatter{})
		instance.SetLevel(logrus.TraceLevel)
	})

	return &LoggerProvider{
		logger: instance,
	}
}

// GetLogger returns a logger with service context
func (lp *LoggerProvider) GetLogger(serviceName string) *logrus.Entry {
	return lp.logger.WithField("service", serviceName)
}

// GetLoggerWithFields returns a logger with service context and additional fields
func (lp *LoggerProvider) GetLoggerWithFields(serviceName string, fields logrus.Fields) *logrus.Entry {
	return lp.logger.WithFields(fields).WithField("service", serviceName)
}

// Global helper function to get logger with request ID from context
func GetLoggerWithRequestID(ctx context.Context) *logrus.Entry {
	requestID := middleware.GetRequestID(ctx)

	// Ensure instance is initialized
	if instance == nil {
		provider := NewLoggerProvider()
		instance = provider.logger
	}

	return instance.WithFields(logrus.Fields{
		"request_id": requestID,
		"service":    "ErrorHandler",
	})
}
