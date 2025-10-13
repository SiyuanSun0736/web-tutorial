package logic

import (
	"go_web/models"

	"github.com/jinzhu/gorm"
)

// DeleteTodo 删除Todo项
func DeleteTodo(db *gorm.DB, id uint) error {
	return db.Delete(&models.Todo{}, id).Error
}
