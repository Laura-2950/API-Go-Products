package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	StorageDB *sql.DB
   )

func init(){
	dataSource := "root:abcd1234@tcp(localhost:3306)/storage"
	// Open inicia un pool de conexiones. Solo abrir una vez
	var err error
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
	panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
	panic(err)
	}
	log.Println("database Configured")
}
   