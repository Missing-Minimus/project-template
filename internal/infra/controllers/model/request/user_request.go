package request

type CreateUserRequest struct {
	Username string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// temos que deixar as entidades assim pra podermos dar unmarshal do JSON direto nela
