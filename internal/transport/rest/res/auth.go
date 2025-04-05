package res

type SignUpRes struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

type SignInRes struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
