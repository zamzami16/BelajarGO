package service

type ValidationService interface {
	Validate(i any) error
}
