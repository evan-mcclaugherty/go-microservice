package app

import (
	"github.com/evan-mcclaugherty/go-microservice/introduction/webserver/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
