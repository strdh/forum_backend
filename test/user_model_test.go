package test

import (
    "errors"
    "testing"
    "time"
    "xyzforum/config"
    "xyzforum/models"
    "github.com/google/uuid"
    "github.com/stretchr/testify/assert"
)

func init() {
    config.InitializeTestDB()
}

func TestUserCreate(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE users")
    if err != nil {
        t.Error(err)
    }

    userModel := models.UserModel{}

    userUUID := uuid.New()
    binaryUUID, _ := userUUID.MarshalBinary()

    user := models.User{
        UserUUID: binaryUUID,
        Username: "testusername",
        Email: "testmail@mail.com",
        Password: "testpassword",
        Avatar: "testavatar",
        Status: 1,
        Created: time.Now().Unix(),
    }

    id, err := userModel.Create(user)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, int64(1), id)
}

func TestUserByUsername(t *testing.T) {
    userModel := models.UserModel{}

    _, err := config.DB.Exec("TRUNCATE TABLE users")
    if err != nil {
        t.Error(err)
    }

    userUUID := uuid.New()
    binaryUUID, _ := userUUID.MarshalBinary()

    user := models.User{
        UserUUID: binaryUUID,
        Username: "testusername",
        Email: "testmail@mail.com",
        Password: "testpassword",
        Created: time.Now().Unix(),
        Avatar: "testavatar",
        Status: 1,
    }

    _, err = config.DB.Exec("INSERT INTO users (user_uuid, username, email, password, created, avatar, status) VALUES(?, ?, ?, ?, ?, ?, ?)", user.UserUUID, user.Username, user.Email, user.Password, user.Created, user.Avatar, user.Status)
    if err != nil {
        t.Error(err)
    }

    result, err := userModel.ByUsername("testusername")
    if err != nil {
        t.Error(err)
    }    

    assert.Equal(t, user.Username, result.Username)
    assert.Equal(t, user.Email, result.Email)
    assert.Equal(t, user.Password, result.Password)
    assert.Equal(t, user.Avatar, result.Avatar)
    assert.Equal(t, user.Status, result.Status)
    assert.Equal(t, user.Created, result.Created)
}

func TestUserByUsername2(t *testing.T) {
    userModel := models.UserModel{}

    _, err := config.DB.Exec("TRUNCATE TABLE users")
    if err != nil {
        t.Error(err)
    }

    userUUID := uuid.New()
    binaryUUID, _ := userUUID.MarshalBinary()

    user := models.User{
        UserUUID: binaryUUID,
        Username: "testusername",
        Email: "testmail@mail.com",
        Password: "testpassword",
        Created: time.Now().Unix(),
        Avatar: "testavatar",
        Status: 1,
    }

    _, err = config.DB.Exec("INSERT INTO users (user_uuid, username, email, password, created, avatar, status) VALUES(?, ?, ?, ?, ?, ?, ?)", user.UserUUID, user.Username, user.Email, user.Password, user.Created, user.Avatar, user.Status)
    if err != nil {
        t.Error(err)
    }

    result, err := userModel.ByUsername("wrongusername")

    assert.Equal(t, models.User{}, result)
    assert.Equal(t, err, errors.New("User not found"))
}

func TestUserUpdate(t *testing.T) {
    userModel := models.UserModel{}

    _, err := config.DB.Exec("TRUNCATE TABLE users")
    if err != nil {
        t.Error(err)
    }

    userUUID := uuid.New()
    binaryUUID, _ := userUUID.MarshalBinary()

    user := models.User{
        UserUUID: binaryUUID,
        Username: "testusername",
        Email: "testmail@mail.com",
        Password: "testpassword",
        Created: time.Now().Unix(),
        Avatar: "testavatar",
        Status: 1,
    }

    _, err = config.DB.Exec("INSERT INTO users (user_uuid, username, email, password, created, avatar, status) VALUES(?, ?, ?, ?, ?, ?, ?)", user.UserUUID, user.Username, user.Email, user.Password, user.Created, user.Avatar, user.Status)
    if err != nil {
        t.Error(err)
    }

    userUpdated := models.User{
        Username: "testusernameupdate",
        Email: "testmailupdate@mail.com",
        Password: "testpasswordupdate",
        Avatar: "testavatarupdate",
    }

    result, err := userModel.Update(userUpdated, 1)
    if err != nil {
        t.Error(err)
    }

    dataUpdated, err := userModel.ByUsername("testusernameupdate")

    assert.Equal(t, int64(1), result)
    assert.Equal(t, userUpdated.Username, dataUpdated.Username)
    assert.Equal(t, userUpdated.Email, dataUpdated.Email)
    assert.Equal(t, userUpdated.Password, dataUpdated.Password)
    assert.Equal(t, userUpdated.Avatar, dataUpdated.Avatar)
}

func TestUserDelete(t *testing.T) {
    userModel := models.UserModel{}

    _, err := config.DB.Exec("TRUNCATE TABLE users")
    if err != nil {
        t.Error(err)
    }

    userUUID := uuid.New()
    binaryUUID, _ := userUUID.MarshalBinary()

    user := models.User{
        UserUUID: binaryUUID,
        Username: "testusername",
        Email: "testmail@mail.com",
        Password: "testpassword",
        Created: time.Now().Unix(),
        Avatar: "testavatar",
        Status: 1,
    }

    _, err = config.DB.Exec("INSERT INTO users (user_uuid, username, email, password, created, avatar, status) VALUES(?, ?, ?, ?, ?, ?, ?)", user.UserUUID, user.Username, user.Email, user.Password, user.Created, user.Avatar, user.Status)
    if err != nil {
        t.Error(err)
    }

    result, err := userModel.Delete(1)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, int64(1), result)
}