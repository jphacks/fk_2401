package mysql

import (
	"database/sql"
	"fmt"
	"log"

	// MySQL driver for database/sql
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	cfg := NewConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("cannot open DB: %v", err)

		return nil, err
	}

	return db, nil
}
