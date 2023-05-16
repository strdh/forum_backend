package handlers

import (
    "time"
    "strconv"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "xyzforum/utils"
    "xyzforum/models"
    "xyzforum/validators"
    "github.com/gorilla/mux"
)

type MessageHandler struct {
    MessageModel *models.MessageModel
    MessageValidator *validators.MessageValidator
}

func (handler *MessageHandler) ByIdForum(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        utils.WriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed", nil)
        return
    }

    id := mux.Vars(r)["id_forum"]
    if id == "" {
        utils.WriteResponse(w, r, http.StatusBadRequest, "Bad request", nil)
        return
    }

    param, _ := strconv.Atoi(id)
    messages := handler.MessageModel.ByIdForum(param)
    utils.WriteResponse(w, r, http.StatusOK, "success", messages)
}

func (handler *MessageHandler) Create(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        utils.WriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed", nil)
        return
    }

    id := mux.Vars(r)["id_forum"]
    if id == "" {
        utils.WriteResponse(w, r, http.StatusBadRequest, "Bad request", nil)
        return
    }

    param, _ := strconv.Atoi(id)
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }
    defer r.Body.Close()

    messageRequest := validators.MessageRequest{}
    err = json.Unmarshal(body, &messageRequest)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    isValid, messages := handler.MessageValidator.ValidateMessage(messageRequest)
    if !isValid {
        utils.WriteResponse(w, r, http.StatusBadRequest, "Bad request", messages)
        return
    }

    //fill the data

    idUser := r.Context().Value("userId").(float64)
    finalIdUser := int(idUser)

    message := models.Message{
        IdForum: param,
        IdUser: finalIdUser,
        Message: messageRequest.Message,
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
    }

    _, err = handler.MessageModel.Create(message)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    utils.WriteResponse(w, r, http.StatusOK, "success", message)
}