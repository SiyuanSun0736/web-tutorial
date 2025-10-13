package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go_web/logic"
)

func GetAllTodosHandler(c *gin.Context) {
	todos, err := logic.GetAllTodos(DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 转换响应格式以确保前端能正确解析
	var response []gin.H
	for _, todo := range todos {
		response = append(response, gin.H{
			"id":        todo.ID,
			"title":     todo.Title,
			"completed": todo.Completed,
			"status":    todo.Completed, // 添加status字段以匹配前端期望
		})
	}
	c.JSON(http.StatusOK, response)
}
