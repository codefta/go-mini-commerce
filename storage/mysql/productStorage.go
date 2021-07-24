package mysql_storage

import (
	"fmt"
	"time"

	"github.com/fathisiddiqi/go-mini-commerce/models"
)

func (s *MysqlStorage) GetProducts(limit int) ([]models.Product, error) {
	queryGet := fmt.Sprintf("SELECT id, product_name, description, price, created_at, updated_at FROM products ORDER BY created_at DESC LIMIT %d", limit)

	rows, err := s.db.Query(queryGet)
	if err != nil {
		return nil, fmt.Errorf("unable to execute query due to: %v", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.ID, &product.ProductName, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			continue
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("unable to iterate rows due to: %v", err)
	}

	return products, nil
}

func (s *MysqlStorage) GetProductById(id int) (*models.Product, error) {
	var product models.Product

	query := `SELECT id, product_name, description, price, created_at, updated_at FROM products where id=?`

	err := s.db.QueryRow(query, id).Scan(&product.ID, &product.ProductName, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("unable to get data due to: %v", err)
	}

	return &product, nil
}

func (s *MysqlStorage) CheckProductIfExist(id int) (int, error) {
	query := `SELECT id FROM products where id = ?`

	err := s.db.QueryRow(query, id).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("data not found")
	}

	return id, nil
}

func (s *MysqlStorage) PostProduct(product models.Product) (*models.Product, error) {
	timeNow := time.Now().Unix()

	query := `INSERT INTO products(product_name, description, price, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	result, err := s.db.Exec(query, product.ProductName, product.Description, product.Price, timeNow, timeNow)
	if err != nil {
		return nil, fmt.Errorf("unable to insert data due to: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("unable to get last inserted id due to: %v", err)
	}

	product.ID = int(id)
	product.CreatedAt = int(timeNow)
	product.UpdatedAt = int(timeNow)

	return &product, nil
}

func (s *MysqlStorage) UpdateProduct(id int, product models.Product) (*models.Product, error) {
	timeNow := time.Now().Unix()

	query := `UPDATE products SET product_name=?, description=?, price=?, updated_at=? WHERE id=?`
	_, err := s.db.Exec(query, product.ProductName, product.Description, product.Price, timeNow, id)
	if err != nil {
		return nil, fmt.Errorf("unable to update data due to: %v", err)
	}

	query = `SELECT id, product_name, description, price, created_at, updated_at FROM products where id=?`

	err = s.db.QueryRow(query, id).Scan(&product.ID, &product.ProductName, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("unable to get data due to: %v", err)
	}

	product.ID = id
	product.UpdatedAt = int(timeNow)

	return &product, nil
}

func (s *MysqlStorage) DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE id=?`
	
	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("unable to delete data due to: %v", err)
	}

	return nil
}