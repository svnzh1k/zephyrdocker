package service

import (
	"database/sql"
	"encoding/json"
	"zephyr-api-mod/internal/models"
)

func CreateCategory(name string, parentId int) error {
	var stmt *sql.Stmt
	var err error

	if parentId == 0 {
		stmt, err = Database.Prepare("INSERT INTO categories (name) VALUES ($1)")
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(name)
		return err
	}

	row := Database.QueryRow("SELECT id FROM categories WHERE id = $1", parentId)
	var id int
	err = row.Scan(&id)
	if err != nil {
		return err
	}

	stmt, err = Database.Prepare("INSERT INTO categories (name, parent_id) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, parentId)
	return err
}

func UpdateCategory(id int, newName string) error {
	stmt, err := Database.Prepare("UPDATE categories SET name = $1 WHERE id = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(newName, id)
	return err
}

func RemoveCategory(id int) error {
	stmt, err := Database.Prepare("DELETE FROM categories WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}

func GetCategories(parentId int) ([]byte, error) {
	var stmt *sql.Stmt
	var err error
	var rows *sql.Rows
	var data []byte

	if parentId == 0 {
		stmt, err = Database.Prepare("SELECT id, name FROM categories WHERE parent_id IS NULL")
		if err != nil {
			return nil, err
		}
		rows, err = stmt.Query()
		if err != nil {
			return nil, err
		}
	} else {
		stmt, err = Database.Prepare("SELECT id, name FROM categories WHERE parent_id = $1")
		if err != nil {
			return nil, err
		}
		rows, err = stmt.Query(parentId)
		if err != nil {
			return nil, err
		}
	}

	var category models.CategoryDTO
	var arr []models.CategoryDTO
	for rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			return nil, err
		}
		arr = append(arr, category)
	}
	data, err = json.Marshal(arr)
	if err != nil {
		return nil, err
	}
	return data, nil
}
