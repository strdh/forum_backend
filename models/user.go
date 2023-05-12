package models

import (
    "log"
    "errors"
    "xyzforum/config"
)

type User struct {
    Id int `json:"id,omitempty"`
    UserUUID []byte `json:"user_uuid,omitempty"`
    Username string `json:"username,omitempty"`
    Email string `json:"email,omitempty"`
    Password string `json:"password,omitempty"`
    Created int64 `json:"created,omitempty"`
    Avatar string `json:"avatar,omitempty"`
    Status int `json:"status,omitempty"`
}

type UserModel struct {}

func (userModel *UserModel) CreateUser(user User) (int64, error) {
    result, err := config.DB.Exec("INSERT INTO users (user_uuid, username, email, password, avatar, status, created) VALUES(?, ?, ?, ?, ?, ?, ?)", user.UserUUID, user.Username, user.Email, user.Password, user.Created, user.Avatar, user.Status)
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
    rows, err := config.DB.Query("SELECT * FROM users WHERE username = ?", username)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    if rows.Next() {
        err := rows.Scan(&user.Id, &user.UserUUID, &user.Username, &user.Email, &user.Password, &user.Created, &user.Avatar, &user.Status)
        if err != nil {
            log.Println(err)
        }

        return user, nil
    }

    return user, errors.New("User not found")
}

func (userModel *UserModel) UpdateUser(user User, id int) (int64, error) {
    result, err := config.DB.Exec("UPDATE users SET username = ?, email = ?, password = ?, avatar = ? WHERE id = ?", user.Username, user.Email, user.Password, user.Avatar, id)
    if err != nil {
        log.Println(err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return 0, err
    }

    return rowsAffected, nil
}

func (userModel *UserModel) DeleteUser(id int) (int64, error) {
    result, err := config.DB.Exec("DELETE FROM users WHERE id = ?", id)
    if err != nil {
        log.Println(err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return 0, err
    }

    return rowsAffected, nil
}