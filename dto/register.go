package dto

// This is like the input!
type RegisterRequestBody struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
