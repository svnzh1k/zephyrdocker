package models

type UserDTO struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type CategoryDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Password string `json:"password"`
	Code     int    `json:"code"`
}

type Category struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Parent_id int    `json:"parent_id"`
}

type Product struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	InStock int    `json:"in_stock"`
	Unit    string `json:"unit"`
}

type Food struct {
	Id           int    `json:"id"`
	Category     int    `json:"category_id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Max_quantity int    `json:"max_quantity"`
}
