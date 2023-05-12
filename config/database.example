package config

import (
    "database/sql"
    "fmt"
    "log"
    "github.com/go-sql-driver/mysql"
)

var DB *sql.DB 

func InitializeDB() {
    dbConfig := mysql.Config{
        User: "dbusername",
        Passwd: "dbpassword",
        Net: "tcp",
        Addr: "127.0.0.1:3306",
        DBName: "dbname",
        AllowNativePasswords: true,
    }

    var err error
    DB, err = sql.Open("mysql", dbConfig.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := DB.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }

    fmt.Println("Database is Connected")
}