package inventory

import (
	"time"
	"uuid"
)

type Product struct {
	Id 					uuid.UUID `json:"id"`
	Name 				string		`json:"name"`
	Description string		`json:"description"`
	Image       string		`json:"image"`
	Price       float64		 `json:"price"`
	Quantity 		int				`json:"quantity"`
	CreatedAt 	time.Time `json:"createdAt"`
	UpdatedAt 	time.Time `json:"updatedAt"`
}