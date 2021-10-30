package models

type Product struct {
	Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
}
