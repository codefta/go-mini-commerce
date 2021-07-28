package models

type Product struct {
	ID 			int			`json:"id"`
	ProductName string		`json:"product_name"`
	Description string		`json:"description"`
	Price 		int			`json:"price"`
	Categories	[]string	`json:"categories"`
	CreatedAt 	int			`json:"created_at"`
	UpdatedAt	int			`json:"updated_at"`
}

type ProductForm struct {
	ProductName		string 		`json:"product_name" form:"required"`
	Description		string		`json:"description"`
	Price			int			`json:"price" form:"numeric"`
	Categories		[]string 	`json:"categories"`
}