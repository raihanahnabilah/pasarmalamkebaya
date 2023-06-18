package dto

// This is like the input!
type RegisterRequestBody struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterEmailVerification struct {
	Name             string
	Subject          string
	Email            string
	VerificationCode string
}
