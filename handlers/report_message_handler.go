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

type ReportMessageHandler struct {
    ReportMessageModel *models.ReportMessageModel
    ReportMessageValidator *validators.ReportMessageValidator
}

func (handler *ReportMessageHandler) Create(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        utils.WriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed", nil)
        return 
    }

    idForum := mux.Vars(r)["id_forum"]
    if idForum == "" {
        utils.WriteResponse(w, r, http.StatusBadRequest, "Bad request", nil)
        return
    }
    finalIdForum, _ := strconv.Atoi(idForum)

    idMessage := mux.Vars(r)["id"]
    if idMessage == "" {
        utils.WriteResponse(w, r, http.StatusBadRequest, "Bad request", nil)
        return
    }
    finalIdMessage, _ := strconv.Atoi(idMessage)

    idReporter := r.Context().Value("userId").(float64)
    finalIdReporter := int(idReporter)

    idForumReal, idOwner := handler.ReportMessageModel.GetIdFO(finalIdMessage)
    if idForumReal != finalIdForum || idOwner == finalIdReporter  {
        utils.WriteResponse(w, r, http.StatusForbidden, "Forbidden", nil)
        return
    }

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }
    defer r.Body.Close()

    reportMessageRequest := validators.ReportMessageRequest{}
    err = json.Unmarshal(body, &reportMessageRequest)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    isValid, messages := handler.ReportMessageValidator.Validate(reportMessageRequest)
    if !isValid {
        utils.WriteResponse(w, r, http.StatusBadRequest, "Bad request", messages)
        return
    }

    reportMessage := models.ReportMessage{
        IdMessage: finalIdMessage,
        IdOwner: idOwner,
        IdReporter: finalIdReporter,
        Problem: reportMessageRequest.Problem,
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
        Status: 0,
    }

    _, err = handler.ReportMessageModel.Create(reportMessage)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    utils.WriteResponse(w, r, http.StatusCreated, "success", reportMessage)
}