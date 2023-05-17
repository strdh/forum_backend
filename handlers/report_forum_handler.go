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

type ReportForumHandler struct {
    ReportForumModel *models.ReportForumModel
    ReportForumValidator *validators.ReportForumValidator
}

func (handler *ReportForumHandler) Create(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        utils.WriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed", nil)
        return
    }

    idForum := mux.Vars(r)["id"]
    if idForum == "" {
        utils.WriteResponse(w, r, http.StatusBadRequest, "Bad request", nil)
        return
    }
    finalIdForum, _ := strconv.Atoi(idForum)

    idOwner := handler.ReportForumModel.GetIdOwner(finalIdForum)
    if idOwner == 0 {
        utils.WriteResponse(w, r, http.StatusBadRequest, "Bad request", nil)
        return
    }

    idReporter := r.Context().Value("userId").(float64)
    finalIdReporter := int(idReporter)

    if finalIdReporter == idOwner {
        utils.WriteResponse(w, r, http.StatusForbidden, "Foribidden", nil)
        return
    }

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    reportForumRequest := validators.ReportForumRequest{}
    err = json.Unmarshal(body, &reportForumRequest)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    isValid, messages := handler.ReportForumValidator.Validate(reportForumRequest)
    if !isValid {
        utils.WriteResponse(w, r, http.StatusBadRequest, "Bad request", messages)
        return
    }

    reportForum := models.ReportForum{
        IdForum: finalIdForum,
        IdOwner: idOwner,
        IdReporter: finalIdReporter,
        Problem: reportForumRequest.Problem,
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
        Status: 0,
    }

    _, err = handler.ReportForumModel.Create(reportForum)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    utils.WriteResponse(w, r, http.StatusOK, "success", reportForum)
}