package response

type UserResponse struct {
	Username string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
