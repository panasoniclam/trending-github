package request

type RequestSignUp struct {
	Email string `validate:"required"`
	FullName string  `validate:"required"`
	Password string  `validate:"required"`
}