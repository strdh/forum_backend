package handlers

import (
    "os"
    "time"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "xyzforum/utils"
    "xyzforum/models"
    "xyzforum/validators"
    "github.com/google/uuid"
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

func (handler *AuthHandler) CheckJWT(w http.ResponseWriter, r *http.Request) {
    
}