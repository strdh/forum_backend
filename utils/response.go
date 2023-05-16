package utils

import (
    "encoding/json"
    "net/http"
)

type Response struct {
    Status int `json:"status"`
    Message string `json:"message"`
    Data interface{} `json:"data"`
}

func WriteResponse(w http.ResponseWriter, r *http.Request, status int, message string, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    response := Response{
        Status: status,
        Message: message,
        Data: data,
    }

    jsonData, err := json.Marshal(response)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.WriteHeader(status)
    w.Write(jsonData)
}