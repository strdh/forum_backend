package config

import (
    "os"
    "log"
    "database/sql"
    "github.com/go-sql-driver/mysql"
)

var DB *sql.DB 

func InitializeDB() {
    dbConfig := mysql.Config{
        User: os.Getenv("DB_USERNAME"),
        Passwd: os.Getenv("DB_PASSWORD"),
        Net: "tcp",
        Addr: os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT"),
        DBName: os.Getenv("DB_NAME"),
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

    log.Println("Database is Connected")
}

func InitializeTestDB() {
    dbConfig := mysql.Config{
        User: os.Getenv("DB_USERNAME"),
        Passwd: os.Getenv("DB_PASSWORD"),
        Net: "tcp",
        Addr: os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT"),
        DBName: os.Getenv("DB_TEST_NAME"),
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

    log.Println("Database is Connected")
}