package auth

type SignInInputSerializer struct {
	Email    string `json:"email,omitempty" binding:"email"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username,omitempty"`
}
