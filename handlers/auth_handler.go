package handlers

import (
    "os"
    "time"
    "strconv"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "xyzforum/utils"
    "xyzforum/models"
    "xyzforum/validators"
    "github.com/google/uuid"
    "github.com/gorilla/mux"
    "golang.org/x/crypto/bcrypt"
)

type loginResponse struct {
    Token string `json:"token"`
}

type AuthHandler struct {
    UserModel *models.UserModel
    AuthValidator *validators.AuthValidator
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

    registerRequest := validators.RegisterRequest{}
    err = json.Unmarshal(body, &registerRequest)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    isValid, messages := handler.AuthValidator.ValidateRegister(registerRequest)
    if !isValid {
        utils.WriteResponse(w, r, http.StatusBadRequest, "Bad request", messages)
        return
    }

    //fill the data
    userUUID := uuid.New()
    binaryUUID, _ := userUUID.MarshalBinary()
    password, err := hashPassword(registerRequest.Password)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    user := models.User{
        UserUUID: binaryUUID,
        Username: registerRequest.Username,
        Email: registerRequest.Email,
        Password: password,
        Created: time.Now().Unix(),
        Avatar: "default.png",
        Status: 1,
    }

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

    loginRequest := validators.LoginRequest{}
    err = json.Unmarshal(body, &loginRequest)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    isValid, messages := handler.AuthValidator.ValidateLogin(loginRequest)
    if !isValid {
        utils.WriteResponse(w, r, http.StatusBadRequest, "Bad request", messages)
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

    jwtToken := utils.GenerateToken(user, os.Getenv("JWT_KEY"))

    loginResponse := loginResponse{
        Token: jwtToken,
    }

    utils.WriteResponse(w, r, http.StatusOK, "Login success", loginResponse)
}

func (handler *AuthHandler) Profile(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        utils.WriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed", nil)
        return
    }

    var user models.User
    var forums []models.Forum
    var messages []models.Message

    username := mux.Vars(r)["username"]
    user, err := handler.UserModel.ByUsername(username)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusNotFound, "User not found", nil)
        return
    }

    forums = handler.UserModel.UserForums(user.Id)
    messages = handler.UserModel.UserMessages(user.Id)

    userProfile := models.UserProfile{
        Id: user.Id,
        Username: user.Username,
        Email: user.Email,
        Avatar: user.Avatar,
        Created: user.Created,
        Status: user.Status,
        Forums: forums,
        Messages: messages,
    }

    utils.WriteResponse(w, r, http.StatusOK, "Success", userProfile)
}

func (handler *AuthHandler) NextMsg(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        utils.WriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed", nil)
        return
    }

    var messages []models.Message

    username := mux.Vars(r)["username"]
    idMsg := mux.Vars(r)["id_msg"]
    finalIdMsg, _ := strconv.Atoi(idMsg)

    id, err := handler.UserModel.GetId(username)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusNotFound, "User not found", nil)
        return
    }

    messages = handler.UserModel.UserMessagesN(id, finalIdMsg)

    utils.WriteResponse(w, r, http.StatusOK, "Success", messages)
}

func (handler *AuthHandler) NextForum(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        utils.WriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed", nil)
        return
    }

    var forums []models.Forum
    username := mux.Vars(r)["username"]
    created := mux.Vars(r)["created"]
    finalCreated, _ := strconv.Atoi(created)

    id, err := handler.UserModel.GetId(username)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusNotFound, "User not found", nil)
        return
    }

    forums = handler.UserModel.UserForumsN(id, finalCreated)

    utils.WriteResponse(w, r, http.StatusOK, "Success", forums)
}