package pg

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

func NewDB(connStr string) (*DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

func main() {
	// 连接参数（直接替换为你的公网IP和密码）
	connStr := "host=43.135.156.166 port=5469 user=root password=F04oEHogKr4uoaQU9xZA dbname=yunfan sslmode=disable"

	// 尝试连接
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer db.Close()

	// Ping测试
	err = db.Ping()
	if err != nil {
		log.Fatal("Ping失败:", err)
	}

	fmt.Println("✅ PostgreSQL 连接成功！")
}
