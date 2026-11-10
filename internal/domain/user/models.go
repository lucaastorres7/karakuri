package user

import (
	"time"
	"uuid"
)

type User struct {
	Id          uuid.UUID 		`json:"id"`
	Name        string 		`json:"name"`
	Email       string 		`json:"email"`
	Password    string 		`json:"-"`
	Gender      string 		`json:"gender"`
	CreatedAt   time.Time `json:"createdAt"`
}