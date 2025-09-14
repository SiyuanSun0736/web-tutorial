package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/SiyuanSun0736/web-tutorial/common"
	"github.com/SiyuanSun0736/web-tutorial/controller"
	"github.com/SiyuanSun0736/web-tutorial/middleware"
	_ "github.com/lib/pq"

	_ "net/http/pprof"
)

func init() {
	var err error
	// 数据库连接配置 - 请根据实际情况修改这些值
	host := "localhost"
	port := 5432
	user := "test"
	password := "123"
	dbname := "testdb"

	conStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	common.Db, err = sql.Open("postgres", conStr)
	if err != nil {
		log.Fatalln(err.Error())
	}

	ctx := context.Background()
	err = common.Db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Database connected")
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: &middleware.BasicAuthMiddleware{},
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./wwwroot"))))
	controller.RegisterRoutes()

	log.Println("Server starting...")
	go http.ListenAndServe(":8000", nil)
	server.ListenAndServe()

}
