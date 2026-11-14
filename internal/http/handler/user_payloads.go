package handler

type RegisterUserPayload struct {
	Name        string 		`json:"name" validate:"required"`
	Email       string 		`json:"email" validate:"required,email"`
	Password    string 		`json:"password" validate:"required,min=6,max=130"`
	Gender      string 		`json:"gender"`
}

type LoginUserPayload struct {
	Email       string 		`json:"email" validate:"required,email"`
	Password    string 		`json:"password" validate:"required"`
}