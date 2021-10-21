package repository

import (
	"todo/models"

	"gorm.io/gorm"
)

type todoRepositoryDB struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) todoRepositoryDB {
	return todoRepositoryDB{db: db}
}

type ITodoRepository interface {
	GetAll() ([]models.Todo, error)
}

func (r todoRepositoryDB) GetAll() ([]models.Todo, error) {
	todo := make([]models.Todo, 0)
	err := r.db.Table(models.TableName.Todo).Find(&todo).Error
	return todo, err
}
