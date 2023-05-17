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

func (messageModel *MessageModel) IsOwned(id int, idUser int, idForum int) bool {
    var tempId int

    err := config.DB.QueryRow("SELECT id_user FROM forum_messages WHERE id = ? AND id_forum = ?", id, idForum).Scan(&tempId)
    if err != nil {
        log.Println(err)
        return false
    }

    return tempId == idUser
}

func (messageModel *MessageModel) Create(message Message) (int64, error) {
    //begin the transaction
    tx, err := config.DB.Begin()
    if err != nil {
        log.Println(err)
    }

    result, err := tx.Exec("INSERT INTO forum_messages (id_forum, id_user, message, created, updated) VALUES (?, ?, ?, ?, ?)", message.IdForum, message.IdUser, message.Message, message.Created, message.Updated)
    if err != nil {
        tx.Rollback()
        log.Println(err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    _, err = tx.Exec("UPDATE forums SET messages = messages + 1 WHERE id = ?", message.IdForum)
    if err != nil {
        tx.Rollback()
        log.Println(err)
    }

    err = tx.Commit()
    if err != nil {
        log.Println(err)
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

func (messageModel *MessageModel) Delete(id int, idForum int) (int64, error) {
    //begin transaction
    tx, err := config.DB.Begin()
    if err != nil {
        log.Println(err)
    }

    result, err := tx.Exec("DELETE FROM forum_messages WHERE id = ?", id)
    if err != nil {
        tx.Rollback()
        log.Println(err)
    }

    _, err = tx.Exec("UPDATE forums SET messages = messages - 1 WHERE id = ?", idForum)
    if err != nil {
        tx.Rollback()
        log.Println(err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        tx.Rollback()
        return 0, err
    }

    err = tx.Commit()
    if err != nil {
        log.Println(err)
    }

    return rowsAffected, nil
}