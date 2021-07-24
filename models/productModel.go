package models

type Product struct {
	ID 			int		`json:"id"`
	ProductName string	`json:"product_name"`
	Description string	`json:"description"`
	Price 		int		`json:"price"`
	CreatedAt 	int		`json:"created_at"`
	UpdatedAt	int		`json:"updated_at"`
}