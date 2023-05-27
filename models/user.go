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

type UserProfile struct {
    Id int `json:"id"`
    Username string `json:"username"`
    Email string `json:"email"`
    Avatar string `json:"avatar"`
    Created int64 `json:"created"`
    Status int `json:"status"`
    Forums []Forum `json:"forums"`
    Messages []Message `json:"messages"`
}

type UserModel struct {}

func (userModel *UserModel) Create(user User) (int64, error) {
    result, err := config.DB.Exec("INSERT INTO users (user_uuid, username, email, password, created, avatar, status) VALUES(?, ?, ?, ?, ?, ?, ?)", user.UserUUID, user.Username, user.Email, user.Password, user.Created, user.Avatar, user.Status)
    if err != nil {
        log.Println(err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return id, nil
}

func (userModel *UserModel) ByUsername(username string) (User, error) {
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

func (userModel *UserModel) Update(user User, id int) (int64, error) {
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

func (userModel *UserModel) Delete(id int) (int64, error) {
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

func (userModel *UserModel) UserForums(idUser int) []Forum {
    var forums []Forum
    var temp Forum

    rows, err := config.DB.Query("SELECT * FROM forums WHERE id_user = ? ORDER BY created ASC LIMIT 15", idUser)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&temp.Id, &temp.IdUser, &temp.Title, &temp.Slug, &temp.Description, &temp.ActiveUsers, &temp.Messages, &temp.Status, &temp.Created)
        if err != nil {
            log.Println(err)
        }

        forums = append(forums, temp)
    }

    return forums
}

func (userModel *UserModel) UserMessages(idUser int) []Message {
    var messages []Message
    var temp Message

    rows, err := config.DB.Query("SELECT * FROM forum_messages WHERE id_user = ? ORDER BY created ASC LIMIT 15", idUser)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&temp.Id, &temp.IdForum, &temp.IdUser, &temp.Message, &temp.Created, &temp.Updated)
        if err != nil {
            log.Println(err)
        }

        messages = append(messages, temp)
    }

    return messages
}