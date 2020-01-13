package rep

type RequestSignUp struct {
	FullName string  `json:"fullname" validate:"required"`
	Email string      `validate:"required"`
	Password string  `validate:"required"`
}