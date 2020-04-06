package auth_oauth2

type loginPayload struct {
	Email	 string `json:"email" validate:"required,email"`
	Password string	`json:"password" validate:"required"`
}