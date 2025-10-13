package logic

import (
	"go_web/models"

	"github.com/jinzhu/gorm"
)

// UpdateTodo 更新Todo项(支持部分更新)
func UpdateTodo(db *gorm.DB, id uint, todo models.Todo) error {
	updates := map[string]interface{}{
		"completed": todo.Completed,
	}
	if todo.Title != "" {
		updates["title"] = todo.Title
	}
	return db.Model(&models.Todo{}).Where("id = ?", id).Updates(updates).Error
}
