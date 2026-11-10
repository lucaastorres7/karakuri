package handler

type RegisterUserPayload struct {
	Name        string 		`json:"name"`
	Email       string 		`json:"email"`
	Password    string 		`json:"password"`
	Gender      string 		`json:"gender"`
}