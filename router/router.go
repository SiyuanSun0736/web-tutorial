package router

import (
	"github.com/gin-gonic/gin"

	"go_web/controller"
	"go_web/middleware"
)

// SetupRouter 配置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 使用中间件
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.CORS())

	// 静态文件
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	// 首页
	r.GET("/", controller.IndexHandler)

	// API 路由组
	v1 := r.Group("v1")
	{
		v1.POST("/todo", controller.AddTodoHandler)
		v1.GET("/todo", controller.GetAllTodosHandler)
		v1.GET("/todo/:id", controller.GetOneTodoHandler)
		v1.PUT("/todo/:id", controller.UpdateTodoHandler)
		v1.DELETE("/todo/:id", controller.DeleteTodoHandler)
	}

	return r
}
