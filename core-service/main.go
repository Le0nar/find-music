package main

import (
	"net/http"

	"github.com/Le0nar/find-music/core-service/internal/handler"
)

func main() {
	handler := handler.NewHandler()

	router := handler.InitRouter()

	http.ListenAndServe(":3000", router)
}