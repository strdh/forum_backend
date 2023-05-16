package handlers

import (
    "log"
    "time"
    "strconv"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "xyzforum/utils"
    "xyzforum/models"
    "xyzforum/validators"
    "github.com/gorilla/mux"
    "github.com/google/uuid"
)

type ForumHandler struct {
    ForumModel *models.ForumModel
    ForumValidator *validators.ForumValidator
}

func (handler *ForumHandler) Forums(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        utils.WriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed", nil)
        return
    }

    forums := handler.ForumModel.Forums()
    utils.WriteResponse(w, r, http.StatusOK, "Success", forums)
}

func (handler *ForumHandler) Create(w http.ResponseWriter, r *http.Request) {
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

    request := validators.ForumRequest{}
    err = json.Unmarshal(body, &request)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    isValid, messages := handler.ForumValidator.ValidateForum(request)
    if !isValid {
        utils.WriteResponse(w, r, http.StatusBadRequest, "Bad request", messages)
        return
    }

    slug := uuid.New()
    finalSlug, _ := slug.MarshalBinary()

    idUser := r.Context().Value("userId").(float64)
    finalIdUser := int(idUser)

    forum := models.Forum{
        IdUser: finalIdUser,
        Title: request.Title,
        Slug: finalSlug,
        Description: request.Description,
        ActiveUsers: 0,
        Messages: 0,
        Status: 1,
        Created: time.Now().Unix(),
    }

    _, err = handler.ForumModel.Create(forum)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    utils.WriteResponse(w, r, http.StatusOK, "Forum created", forum)
}

func (handler *ForumHandler) ById(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        utils.WriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed", nil)
    }

    id := mux.Vars(r)["id"]
    if id == "" {
        utils.WriteResponse(w, r, http.StatusBadRequest, "Bad request", nil)
        return
    }

    param, _ := strconv.Atoi(id) 
    forum, err := handler.ForumModel.ById(param)
    log.Println(param)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    utils.WriteResponse(w, r, http.StatusOK, "Success", forum)
}

func (handler *ForumHandler) Update(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        utils.WriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed", nil)
        return
    }

    idForum := mux.Vars(r)["id"]
    if idForum == "" {
        utils.WriteResponse(w, r, http.StatusBadRequest, "Bad request", nil)
        return
    }

    param, _ := strconv.Atoi(idForum)
    idUser := r.Context().Value("userId").(float64)
    finalIdUser := int(idUser)

    isOwned := handler.ForumModel.IsOwned(param, finalIdUser)
    if !isOwned {
        utils.WriteResponse(w, r, http.StatusForbidden, "Foribidden", nil)
        return
    }

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    request := validators.ForumRequest{}
    err = json.Unmarshal(body, &request)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    isValid, messages := handler.ForumValidator.ValidateForum(request)
    if !isValid {
        utils.WriteResponse(w, r, http.StatusBadRequest, "Bad request", messages)
        return
    }

    forum := models.Forum{
        Title: request.Title,
        Description: request.Description,
    }

    row, err := handler.ForumModel.Update(forum, param)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    utils.WriteResponse(w, r, http.StatusOK, "Forum updated", row)
}

func (handler *ForumHandler) Delete(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        utils.WriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed", nil)
        return
    }

    idForum := mux.Vars(r)["id"]
    if idForum == "" {
        utils.WriteResponse(w, r, http.StatusBadRequest, "Bad request", nil)
        return
    }

    param, _ := strconv.Atoi(idForum)
    idUser := r.Context().Value("userId").(float64)
    finalIdUser := int(idUser)

    isOwned := handler.ForumModel.IsOwned(param, finalIdUser)
    if !isOwned {
        utils.WriteResponse(w, r, http.StatusForbidden, "Foribidden", nil)
        return
    }

    _, err := handler.ForumModel.Delete(param)
    if err != nil {
        utils.WriteResponse(w, r, http.StatusInternalServerError, "server error", nil)
        return
    }

    utils.WriteResponse(w, r, http.StatusOK, "Forum deleted successfully", nil)
}

// func (handler *ForumHandler) Update(w http.ResponseWriter, r *http.Request) {
//     if r.Method != http.MethodPut {
//         utils.WriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed", nil)
//     }


// }