package repository

import (
	// "github.com/gofiber/fiber/v2"
	"github.com/xXNeonKitsuneXx/toDoList_Backend/entities"
	"gorm.io/gorm"
)

type toDoListRepositoryDB struct {
	db *gorm.DB
}

func NewToDoListRepositoryDB(db *gorm.DB) toDoListRepositoryDB {
	return toDoListRepositoryDB{db: db}
}

func (r toDoListRepositoryDB) GetAll() ([]entities.Todo, error) {
	todolists := []entities.Todo{}
	result := r.db.Find(&todolists)
	if result.Error != nil {
		return nil, result.Error
	}
	return todolists, nil
}
