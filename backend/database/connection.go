package database

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
    var err error
    DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/mini_hospital")
    if err != nil {
        panic(err)
    }

    err = DB.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Database connected!")
}