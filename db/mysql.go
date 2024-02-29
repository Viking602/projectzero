package db

import (
	"context"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"projectzero/ent"
	"projectzero/ent/migrate"
	"time"
)

func Database() *ent.Client {
	godotenv.Load()
	c, err := sql.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	db := c.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	// 初始化 Ent 客户端
	client := ent.NewClient(ent.Driver(c))

	// 执行数据库迁移
	err = client.Schema.Create(context.Background(), migrate.WithDropIndex(true))
	if err != nil {
		panic(err)
	}

	return client

}
