package redis_storage

import (
	"encoding/json"
	"fmt"

	"github.com/fathisiddiqi/go-mini-commerce/models"
)

func (r *RedisStorage) SetAllTempProducts(key string, products []models.Product) error {
	productsJSON, err := json.Marshal(products)
	if err != nil {
		return fmt.Errorf("unable parse to json due to: %v", err)
	}

	err = r.client.Set(key, productsJSON, 0).Err()
	if err != nil {
		return fmt.Errorf("unable to set redis value due to: %v", err)
	}

	return nil
}

func (r *RedisStorage) GetAllTempProducts(key string) ([]models.Product, error) {
	var products []models.Product

	productsRaw, err := r.client.Get(key).Result()
	if err != nil {
		return nil, fmt.Errorf("unable to get redis value due to: %v", err)
	}

	err = json.Unmarshal([]byte(productsRaw), &products)
	if err != nil {
		return nil, fmt.Errorf("unable parse to object due to: %v ", err)
	}

	return products, nil
}

func (r *RedisStorage) SetTempProduct(key string, product models.Product) error {
	productJSON, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("unable parse to json due to: %v", err)
	}

	err = r.client.Set(key, productJSON, 0).Err()
	if err != nil {
		return fmt.Errorf("unable to set redis value due to: %v", err)
	}

	return nil
}

func (r *RedisStorage) GetTempProduct(key string) (*models.Product, error) {
	var product models.Product

	productRaw, err := r.client.Get(key).Result()
	if err != nil {
		return nil, fmt.Errorf("unable to get redis value due to: %v", err)
	}

	err = json.Unmarshal([]byte(productRaw), &product)
	if err != nil {
		return nil, fmt.Errorf("unable parse to object due to: %v ", err)
	}

	return &product, nil
}

func (r *RedisStorage) DeleteAllTempProductsData() {
	var cursor uint64
	searchPattern := "product*"

	iter := r.client.Scan(cursor, searchPattern, 0).Iterator()

	for iter.Next() {
		r.client.Del(iter.Val())
	}
}