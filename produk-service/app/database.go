package app

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/helper"
	_ "github.com/go-sql-driver/mysql"
)

// func New(db *sql.DB) *DB {
// 	return &DB{
// 		db: db,
// 	}
// }

func Init() *sql.DB {

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Some error occured. Err: %s", err)
	// }

	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_DB := os.Getenv("MYSQL_DATABASE")
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")

	fmt.Println("nidzam", MYSQL_DB)
	// dsn := fmt.Sprintf("%s:%s@/%s?parseTime=true", conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_NAME)
	// connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
	// 	conf.DB_HOST, conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_PORT, conf.DB_NAME)

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_DB))
	helper.PanicIfError(err)

	// db.SetMaxIdleConns(5)
	// db.SetMaxOpenConns(20)
	// db.SetConnMaxLifetime(60 * time.Minute)
	// db.SetConnMaxIdleTime(10 * time.Minute)

	// return db

	// db, err := sql.Open("mysql", connString)
	// if err != nil {
	// 	panic(err)
	// }

	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(20)
	// return db
	return db
}
