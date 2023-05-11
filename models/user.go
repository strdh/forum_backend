package models

import (
    "log"
    "errors"
    "time"
    "xyzforum/config"
    "github.com/google/uuid"
)

type User struct {
    Id int `json:"id,omitempty"`
    User_UUID string `json:"user_uuid,omitempty"`
    Username string `json:"username,omitempty"`
    Email string `json:"email,omitempty"`
    Password string `json:"password,omitempty"`
    Avatar string `json:"avatar,omitempty"`
    Status string `json:"status,omitempty"`
    Created int `json:"created,omitempty"`
}

type UserModel struct {}

func (userModel *UserModel) CreateUser(user User) (int64, error) {
    userUUID := uuid.New()

    now := time.Now()
    created := now.Unix()

    result, err := config.DB.Exec("INSERT INTO users (user_uuid, username, email, password, created) VALUES(?, ?, ?, ?, ?)", userUUID, user.Username, user.Email, user.Password, created)
    if err != nil {
        log.Println(err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return id, nil
}

func (userModel *UserModel) GetUserByUsername(username string) (User, error) {
    var user User
    rows, err := config.DB.Query("SELECT username, password FROM users WHERE username = ?", username)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    if rows.Next() {
        err := rows.Scan(&user.Username, &user.Password)
        if err != nil {
            log.Println(err)
        }

        return user, nil
    }

    return user, errors.New("User not found")
}