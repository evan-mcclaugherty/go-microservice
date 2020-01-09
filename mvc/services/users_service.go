package services

import (
	"github.com/evan-mcclaugherty/go-microservice/introduction/webserver/mvc/domain"
	"github.com/evan-mcclaugherty/go-microservice/introduction/webserver/mvc/utils"
)

func GetUser(userId int) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
