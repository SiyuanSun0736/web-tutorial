package logic

import (
	"go_web/models"

	"github.com/jinzhu/gorm"
)

// AddTodo 添加一个新的Todo项并返回创建的Todo
func AddTodo(db *gorm.DB, title string) (models.Todo, error) {
	todo := models.Todo{
		Title:     title,
		Completed: false,
	}
	err := db.Create(&todo).Error
	return todo, err
}
