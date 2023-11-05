package models

type UserModels struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Pw    string `json:"password"`
}
