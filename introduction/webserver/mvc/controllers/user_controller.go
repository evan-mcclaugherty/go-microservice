package controllers

import (
	"github.com/evan-mcclaugherty/go-microservice/introduction/webserver/mvc/services"
	"github.com/evan-mcclaugherty/go-microservice/introduction/webserver/mvc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userIdParam := c.Param("user_id")
	intUserId, err := strconv.Atoi(userIdParam)
	if err != nil {
		appErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, appErr)
		return
	}
	user, appErr := services.GetUser(intUserId)
	if appErr != nil {
		utils.RespondError(c, appErr)
		return
	}
	utils.Respond(c, http.StatusOK, user)
}
