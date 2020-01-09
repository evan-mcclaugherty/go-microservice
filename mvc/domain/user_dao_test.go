package domain

import (
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	const userId int = 10
	user, err := GetUser(userId)
	if user != nil {
		t.Errorf("we were expecting user to be nil with user id %v", userId)
	}
	if err == nil {
		t.Errorf("we were expecting an error with user id %v", userId)
		if err.StatusCode != 404 {
			t.Error("we were expecting a status code of 404")
		}
	}
}

func TestGetUser(t *testing.T) {
	const userId int = 1
	user, err := GetUser(userId)
	if user == nil {
		t.Errorf("we were expecting user with user id %v", userId)
	}
	if err != nil {
		t.Errorf("we were expecting an error with user id %v", userId)
	}
	if user != nil {
		if user.FirstName != "Evan" &&
			user.LastName != "Mack" &&
			user.Email != "e@m.com" &&
			user.Id != userId {
			t.Error()
		}
	}
}
