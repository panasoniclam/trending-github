package rep

type RequestSignUp struct {
	FullName string  `json:"fullname" validate:"required"`
	Email string      `json:"email" validate:"required"`
	Password string  `json:"password" validate:"pwd"`
}