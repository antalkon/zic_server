package req

type SignInReq struct {
	Email    string `json:"phone_number" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}
