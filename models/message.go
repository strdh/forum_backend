package models

import (
    "log"
    "errors"
    "xyzforum/config"
)

type Message struct {
    Id int `json:"id,omitempty"`
    IdForum int `json:"id_forum,omitempty"`
    IdUser int `json:"id_user,omitempty"`
    Message string `json:"message,omitempty"`
    Created int `json:"created,omitempty"`
    Updated int `json:"updated,omitempty"`
}

type MessageModel struct {}

func (messageModel *MessageModel) GetMessageByIdUser(idUser int) []Message {
    var messages []Message
    var temp Message

    rows, err := config.DB.Query("SELECT * FROM messages WHERE id_user = ?", idUser)
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

func (messageModel *MessageModel) GetMessageByIdForum(idForum int) []Message {
    var messages []Message
    var temp Message

    rows, err := config.DB.Query("SELECT * FROM messages WHERE id_forum = ?", idForum)
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

func (messageModel *MessageModel) CreateMessage(message Message) (int64, error) {
    result, err := config.DB.Exec("INSERT INTO messages (id_forum, id_user, message, created, updated) VALUES (?, ?, ?, ?, ?)", message.IdForum, message.IdUser, message.Message, message.Created, message.Updated)
    if err != nil {
        log.Println(err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return id, nil
}

func (messageModel *MessageModel) UpdateMessage(message Message, id int) (Message, error) {
    _, err := config.DB.Exec("UPDATE messages SET message = ?, updated = ? WHERE id = ?", message.Message, message.Updated, id)
    if err != nil {
        log.Println(err)
        return message, errors.New("Update has failed: " + err.Error())
    }

    return message, nil
}

func (messageModel *MessageModel) DeleteMessage(id int) error {
    _, err := config.DB.Exec("DELETE FROM messages WHERE id = ?", id)
    if err != nil {
        log.Println(err)
        return errors.New("Delete has failed: " + err.Error())
    }

    return nil
}