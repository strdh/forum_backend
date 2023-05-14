package models

import (
    "log"
    "xyzforum/config"
)

type Message struct {
    Id int `json:"id,omitempty"`
    IdForum int `json:"id_forum,omitempty"`
    IdUser int `json:"id_user,omitempty"`
    Message string `json:"message,omitempty"`
    Created int64 `json:"created,omitempty"`
    Updated int64 `json:"updated,omitempty"`
}

type MessageModel struct {}

func (messageModel *MessageModel) ByIdUser(idUser int) []Message {
    var messages []Message
    var temp Message

    rows, err := config.DB.Query("SELECT * FROM forum_messages WHERE id_user = ?", idUser)
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

func (messageModel *MessageModel) ByIdForum(idForum int) []Message {
    var messages []Message
    var temp Message

    rows, err := config.DB.Query("SELECT * FROM forum_messages WHERE id_forum = ?", idForum)
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

func (messageModel *MessageModel) Create(message Message) (int64, error) {
    result, err := config.DB.Exec("INSERT INTO forum_messages (id_forum, id_user, message, created, updated) VALUES (?, ?, ?, ?, ?)", message.IdForum, message.IdUser, message.Message, message.Created, message.Updated)
    if err != nil {
        log.Println(err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return id, nil
}

func (messageModel *MessageModel) Update(message Message, id int) (int64, error) {
    result, err := config.DB.Exec("UPDATE forum_messages SET message = ?, updated = ? WHERE id = ?", message.Message, message.Updated, id)
    if err != nil {
        log.Println(err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return 0, err
    }

    return rowsAffected, nil
}

func (messageModel *MessageModel) Delete(id int) (int64, error) {
    result, err := config.DB.Exec("DELETE FROM forum_messages WHERE id = ?", id)
    if err != nil {
        log.Println(err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return 0, err
    }

    return rowsAffected, nil
}