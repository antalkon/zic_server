package models

// запрос авторизации
type Login struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
