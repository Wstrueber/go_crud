package models

// Todo struct (model)
type Todo struct {
	ID          int    `json:"id"`
	Content     string `json:"content"`
	Status      bool   `json:"status"`
	OrderNumber int    `json:"orderNumber"`
}
