package main

import (
    "encoding/json"
    "net/http"
    "bscscan-go/service"

    "github.com/gorilla/mux"
)

type ErrorResponse struct {
    Status  int     `json:"status"`
    Message string  `json:"message"`
}

func (app *application) getContracts(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("Content-Type", "application/json")

    variables := mux.Vars(request)
    address := variables["address"]

    transactions, err := service.RequestContract(app.config.apiKey, address)
    if err != nil {
        app.writeError(response, http.StatusInternalServerError, err.Error())

        return
    }

    jsonResponse, err := json.Marshal(transactions)
    if err != nil {
        app.writeError(response, http.StatusInternalServerError, err.Error())

        return
    }

    response.WriteHeader(http.StatusOK)
    response.Write(jsonResponse)
}

func (app *application) writeError(response http.ResponseWriter, status int, message string) {
    errorResponse := ErrorResponse{
        Status:     status,
        Message:    message,
    }

    jsonResponse, err := json.Marshal(errorResponse)
    if err != nil {
        app.errorLog.Println(err)

        return
    }

    response.WriteHeader(status)
    response.Write(jsonResponse)
}
