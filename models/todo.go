package models

type Todo struct {
	ID   int    `gorm:"primaryKey;column:todo_id" json:"todo_id"`
	Name string `gorm:"column:name" json:"name"`
}

type TodoResponse struct {
	Name string `json:"name"`
}
