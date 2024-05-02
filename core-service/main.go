package main

import (
	"net/http"

	"github.com/Le0nar/find-music/core-service/internal/handler"
	"github.com/Le0nar/find-music/core-service/internal/service"
)

func main() {
	service := service.NewService()
	handler := handler.NewHandler(service)

	router := handler.InitRouter()

	http.ListenAndServe(":3000", router)
}