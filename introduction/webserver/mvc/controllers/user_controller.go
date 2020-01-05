package controllers

import (
	"encoding/json"
	"github.com/evan-mcclaugherty/go-microservice/introduction/webserver/mvc/services"
	"github.com/evan-mcclaugherty/go-microservice/introduction/webserver/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIdParam := r.URL.Query().Get("user_id")
	parsedUserId, err := strconv.Atoi(userIdParam)
	if err != nil {
		appErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		sendErrorResponse(w, appErr)
		return
	}
	user, appErr := services.GetUser(parsedUserId)
	if appErr != nil {
		sendErrorResponse(w, appErr)
		return
	}
	userJson, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/userJson")
	w.Write(userJson)
}

func sendErrorResponse(w http.ResponseWriter, err *utils.ApplicationError) {
	w.Header().Set("Content-Type", "application/errJson")
	w.WriteHeader(err.StatusCode)
	errJson, _ := json.Marshal(err)
	w.Write(errJson)
}
