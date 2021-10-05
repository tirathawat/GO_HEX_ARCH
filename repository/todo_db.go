package repository

import (
	"gorm.io/gorm"
)

type todoRepositoryDB struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) todoRepositoryDB {
	return todoRepositoryDB{db: db}
}

func (r todoRepositoryDB) GetAll() ([]Todo, error) {
	todo := make([]Todo, 0)
	err := r.db.Table(table.Todo).Find(&todo).Error
	return todo, err
}
