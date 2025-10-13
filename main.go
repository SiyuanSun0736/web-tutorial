package main

import (
	"log"

	_ "github.com/lib/pq"

	"go_web/config"
	"go_web/controller"
	"go_web/dao"
	"go_web/router"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 初始化数据库
	db := dao.InitDB()
	if db == nil {
		log.Fatal("Failed to initialize database")
	}
	controller.DB = db
	defer db.Close()

	// 设置路由
	r := router.SetupRouter()

	// 启动服务器
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := r.Run(cfg.Server.Port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
