package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go_web/logic"
)

func DeleteTodoHandler(c *gin.Context) {
	id := uint(0)
	if _, err := fmt.Sscanf(c.Param("id"), "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	if err := logic.DeleteTodo(DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
