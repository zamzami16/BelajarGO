package service

import "github.com/go-playground/validator"

type ValidationServiceImpl struct {
	validator *validator.Validate
}

func NewValidationService() ValidationService {
	return &ValidationServiceImpl{
		validator: validator.New(),
	}
}

func (v *ValidationServiceImpl) Validate(i any) error {
	return v.validator.Struct(i)
}
