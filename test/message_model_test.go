package test

import (
    "time"
    "errors"
    "testing"
    "xyzforum/config"
    "xyzforum/models"
    "github.com/joho/godotenv"
    "github.com/stretchr/testify/assert"
)

func init() {
    err := godotenv.Load("../.env")
    if err != nil {
        panic(err)
    }
    config.InitializeTestDB()
}

func TestMessageByIdUser(t *testing.T) {
    messageModel := models.MessageModel{}

    _, err := config.DB.Exec("TRUNCATE TABLE forum_messages")
    if err != nil {
        t.Error(err)
    }

    message := models.Message{
        IdForum: 1,
        IdUser: 1,
        Message: "testmessage",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
    }

    _, err = config.DB.Exec("INSERT INTO forum_messages (id_forum, id_user, message, created, updated) VALUES (?, ?, ?, ?, ?)", message.IdForum, message.IdUser, message.Message, message.Created, message.Updated)
    if err != nil {
        t.Error(err)
    }

    messages := messageModel.ByIdUser(1)
    if len(messages) != 1 {
        t.Error(errors.New("Message not found"))
    }

    assert.Equal(t, "testmessage", messages[0].Message)
}

func TestMessageByIdForum(t *testing.T) {
    messageModel := models.MessageModel{}

    _, err := config.DB.Exec("TRUNCATE TABLE forum_messages")
    if err != nil {
        t.Error(err)
    }

    message := models.Message{
        IdForum: 1,
        IdUser: 1,
        Message: "testmessage",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
    }

    _, err = config.DB.Exec("INSERT INTO forum_messages (id_forum, id_user, message, created, updated) VALUES (?, ?, ?, ?, ?)", message.IdForum, message.IdUser, message.Message, message.Created, message.Updated)
    if err != nil {
        t.Error(err)
    }

    messages := messageModel.ByIdForum(1)
    if len(messages) != 1 {
        t.Error(errors.New("Message not found"))
    }

    assert.Equal(t, "testmessage", messages[0].Message)
}

func TestMessageCreate(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE forum_messages")
    if err != nil {
        t.Error(err)
    }

    messageModel := models.MessageModel{}

    message := models.Message{
        IdForum: 1,
        IdUser: 1,
        Message: "testmessage",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
    }

    id, err := messageModel.Create(message)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, int64(1), id)
}

func TestMessageUpdate(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE forum_messages")
    if err != nil {
        t.Error(err)
    }

    messageModel := models.MessageModel{}

    message := models.Message{
        IdForum: 1,
        IdUser: 1,
        Message: "testmessage",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
    }

    _, err = config.DB.Exec("INSERT INTO forum_messages (id_forum, id_user, message, created, updated) VALUES (?, ?, ?, ?, ?)", message.IdForum, message.IdUser, message.Message, message.Created, message.Updated)
    if err != nil {
        t.Error(err)
    }

    message.Message = "testmessageupdate"

    result, err := messageModel.Update(message, 1)
    if err != nil {
        t.Error(err)
    }

    updatedMessage := messageModel.ByIdForum(1)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, int64(1), result)
    assert.Equal(t, "testmessageupdate", updatedMessage[0].Message)
}

func TestMessageDelete(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE forum_messages")
    if err != nil {
        t.Error(err)
    }

    messageModel := models.MessageModel{}

    message := models.Message{
        IdForum: 1,
        IdUser: 1,
        Message: "testmessage",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
    }

    _, err = config.DB.Exec("INSERT INTO forum_messages (id_forum, id_user, message, created, updated) VALUES (?, ?, ?, ?, ?)", message.IdForum, message.IdUser, message.Message, message.Created, message.Updated)
    if err != nil {
        t.Error(err)
    }

    result, err := messageModel.Delete(1, 1)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, int64(1), result)
}