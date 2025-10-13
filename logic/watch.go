package logic

import (
	"go_web/models"

	"github.com/jinzhu/gorm"
)

// GetAllTodos 获取所有Todo项
func GetAllTodos(db *gorm.DB) ([]models.Todo, error) {
	var todos []models.Todo
	err := db.Find(&todos).Error
	return todos, err
}

// GetTodoByID 根据ID获取单个Todo项
func GetTodoByID(db *gorm.DB, id uint) (models.Todo, error) {
	var todo models.Todo
	err := db.First(&todo, id).Error
	return todo, err
}
