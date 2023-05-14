package handlers

import (
    // "fmt"
    "time"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "xyzforum/models"
    "xyzforum/utils"
    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type loginResponse struct {
    User models.User `json:"user"`
    Token string `json:"token"`
}

type AuthHandler struct {
    UserModel *models.UserModel
}

func hashPassword(password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hash), nil
}

func comparePassword(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}

func (handler *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        utils.WriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed", nil)
        return
    }

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }
    defer r.Body.Close()

    user := models.User{}
    err = json.Unmarshal(body, &user)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    //fill the data
    userUUID := uuid.New()
    binaryUUID, _ := userUUID.MarshalBinary()

    password, err := hashPassword(user.Password)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    user.UserUUID = binaryUUID
    user.Password = password
    user.Created = time.Now().Unix()
    user.Avatar = "default.png"
    user.Status = 1

    id, err := handler.UserModel.Create(user)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }
    user.Id = int(id)
    user.Password = "*****"

    utils.WriteResponse(w, r, http.StatusOK, "User registered successfully", user)
}

func (handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        utils.WriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed", nil)
        return
    }

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }
    defer r.Body.Close()

    loginRequest := loginRequest{}
    err = json.Unmarshal(body, &loginRequest)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    user, err := handler.UserModel.ByUsername(loginRequest.Username)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusUnauthorized, "username or password invalid", nil)
        return
    }

    if !comparePassword(user.Password, loginRequest.Password) {
        utils.WriteResponse(w, r, http.StatusUnauthorized, "username or password invalid", nil)
        return
    }

    loginResponse := loginResponse{
        User: user,
        Token: "ExampleToken",
    }

    utils.WriteResponse(w, r, http.StatusOK, "Login success", loginResponse)
}

// func (handler *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
    
// }