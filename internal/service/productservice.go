package service

import (
	"encoding/json"
	"zephyr-api-mod/internal/models"
)

func CreateProduct(name string, inStock int, unit string) error {
	stmt, err := Database.Prepare("INSERT INTO products (name, in_stock, unit) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, inStock, unit)
	return err
}

func UpdateProduct(id int, newName string, inStock int, unit string) error {
	stmt, err := Database.Prepare("UPDATE products SET name = $1, in_stock = $2, unit = $3 WHERE id = $4")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(newName, inStock, unit, id)
	return err
}

func RemoveProduct(id int) error {
	stmt, err := Database.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}

func GetProducts() ([]byte, error) {
	rows, err := Database.Query("SELECT id, name, in_stock, unit FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var product models.Product
	var products []models.Product

	for rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.InStock, &product.Unit)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	data, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}
	return data, nil
}
