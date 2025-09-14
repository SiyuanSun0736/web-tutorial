package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

func initDatabase() {
	// 数据库连接配置
	host := "localhost"
	port := 5432
	user := "test"
	password := "123"
	dbname := "testdb"

	conStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("数据库连接测试失败: %v", err)
	}
	log.Println("数据库连接成功")

	// 读取 SQL 文件
	sqlBytes, err := ioutil.ReadFile("init_database.sql")
	if err != nil {
		log.Fatalf("读取 SQL 文件失败: %v", err)
	}

	sqlContent := string(sqlBytes)

	// 分割 SQL 语句（按分号分割，但排除函数定义中的分号）
	statements := splitSQLStatements(sqlContent)

	// 执行每个 SQL 语句
	for i, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" || strings.HasPrefix(stmt, "--") || strings.HasPrefix(stmt, "\\d") || strings.HasPrefix(stmt, "SELECT") {
			continue // 跳过空语句、注释和查询语句
		}

		log.Printf("执行 SQL 语句 %d: %s...", i+1, stmt[:min(50, len(stmt))])
		_, err = db.ExecContext(ctx, stmt)
		if err != nil {
			log.Printf("执行 SQL 语句失败: %v\nSQL: %s", err, stmt)
			continue
		}
		log.Printf("SQL 语句 %d 执行成功", i+1)
	}

	// 验证表是否创建成功
	log.Println("验证表结构...")
	rows, err := db.QueryContext(ctx, "SELECT column_name, data_type FROM information_schema.columns WHERE table_name = 'company' ORDER BY ordinal_position")
	if err != nil {
		log.Printf("查询表结构失败: %v", err)
	} else {
		log.Println("company 表结构:")
		for rows.Next() {
			var columnName, dataType string
			rows.Scan(&columnName, &dataType)
			log.Printf("  %s: %s", columnName, dataType)
		}
		rows.Close()
	}

	// 查看插入的数据
	log.Println("查看插入的数据...")
	rows, err = db.QueryContext(ctx, "SELECT id, name, nickname FROM company")
	if err != nil {
		log.Printf("查询数据失败: %v", err)
	} else {
		log.Println("company 表数据:")
		for rows.Next() {
			var id, name, nickname string
			rows.Scan(&id, &name, &nickname)
			log.Printf("  ID: %s, Name: %s, NickName: %s", id, name, nickname)
		}
		rows.Close()
	}

	log.Println("数据库初始化完成!")
}

func splitSQLStatements(sql string) []string {
	var statements []string
	var current strings.Builder
	inFunction := false

	lines := strings.Split(sql, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// 检查是否进入函数定义
		if strings.Contains(strings.ToUpper(line), "CREATE OR REPLACE FUNCTION") {
			inFunction = true
		}

		// 检查是否结束函数定义
		if inFunction && strings.Contains(line, "$$") && strings.Count(current.String(), "$$") >= 1 {
			inFunction = false
		}

		current.WriteString(line)
		current.WriteString("\n")

		// 如果不在函数定义中，且行以分号结尾，则认为是一个完整的语句
		if !inFunction && strings.HasSuffix(line, ";") {
			statements = append(statements, current.String())
			current.Reset()
		}
	}

	// 添加最后一个语句（如果有的话）
	if current.Len() > 0 {
		statements = append(statements, current.String())
	}

	return statements
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
