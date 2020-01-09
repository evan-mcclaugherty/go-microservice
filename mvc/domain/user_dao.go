package domain

import (
	"fmt"
	"github.com/evan-mcclaugherty/go-microservice/introduction/webserver/mvc/utils"
	"net/http"
)

var (
	users = map[int]*User{
		1: &User{
			Id:        1,
			FirstName: "Evan",
			LastName:  "Mack",
			Email:     "e@m.com",
		},
		2: &User{
			Id:        2,
			FirstName: "John",
			LastName:  "Jingle",
			Email:     "j@j.com",
		},
	}
)

func GetUser(userId int) (*User, *utils.ApplicationError) {
	if user, ok := users[userId]; !ok {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("user %v was not found", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	} else {
		return user, nil
	}
}
