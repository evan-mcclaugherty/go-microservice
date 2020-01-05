package app

import (
	"github.com/evan-mcclaugherty/go-microservice/introduction/webserver/mvc/controllers"
	"log"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
