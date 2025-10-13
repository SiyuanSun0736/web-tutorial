package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go_web/logic"
	"go_web/models"
)

func AddTodoHandler(c *gin.Context) {
	var todo models.Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTodo, err := logic.AddTodo(DB, todo.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTodo)
}
