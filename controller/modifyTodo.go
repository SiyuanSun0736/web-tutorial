package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go_web/logic"
)

func UpdateTodoHandler(c *gin.Context) {
	id := uint(0)
	if _, err := fmt.Sscanf(c.Param("id"), "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var requestBody struct {
		Status bool `json:"status"`
	}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取原有todo
	existingTodo, err := logic.GetTodoByID(DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// 更新完成状态
	existingTodo.Completed = requestBody.Status

	// 更新todo
	if err := logic.UpdateTodo(DB, id, existingTodo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回更新后的todo，确保返回的状态正确
	updatedTodo, err := logic.GetTodoByID(DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":        updatedTodo.ID,
		"title":     updatedTodo.Title,
		"completed": updatedTodo.Completed,
		"status":    updatedTodo.Completed, // 添加status字段以匹配前端期望
	})
}
