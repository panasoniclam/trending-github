package req

type UserRequest struct {
	Email string `validate:"required"`
	FullName string `validate:"required"`
	Pssword string `validate:"required"`
}
