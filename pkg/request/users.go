package request

type RegisterUserInput struct {
	Fullname   string `json:"fullname" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailAvailable struct {
	Email string `json:"email" binding:"required,email"`
}
