package service

import (
	"zephyr-api-mod/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func GetUserByUsername(username string) (*models.User, error) {
	stmt, err := Database.Prepare("SELECT id, username, role, password, code FROM users WHERE username = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res := stmt.QueryRow(username)
	var user models.User
	err = res.Scan(&user.Id, &user.Username, &user.Role, &user.Password, &user.Code)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserById(id int) (*models.User, error) {
	stmt, err := Database.Prepare("SELECT id, username, role, password, code FROM users WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res := stmt.QueryRow(id)
	var user models.User
	err = res.Scan(&user.Id, &user.Username, &user.Role, &user.Password, &user.Code)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetWaiters() (*[]models.UserDTO, error) {
	stmt, err := Database.Prepare("SELECT id, username, role FROM users where role = 'waiter'")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	var arr []models.UserDTO

	for rows.Next() {
		var userDto models.UserDTO
		rows.Scan(&userDto.Id, &userDto.Username, &userDto.Role)
		arr = append(arr, userDto)
	}
	return &arr, err
}

func AddUser(user *models.User) error {
	stmt, err := Database.Prepare("INSERT INTO users (username, password, role, code) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	defer stmt.Close()
	_, err = stmt.Exec(user.Username, user.Password, user.Role, user.Code)
	return err
}

func RemoveUser(id int) error {
	stmt, err := Database.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}
