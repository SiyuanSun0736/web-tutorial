package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go_web/logic"
)

func GetOneTodoHandler(c *gin.Context) {
	id := uint(0)
	if _, err := fmt.Sscanf(c.Param("id"), "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	todo, err := logic.GetTodoByID(DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}
