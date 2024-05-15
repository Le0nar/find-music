package main

import (
	"log"
	"net/http"

	"github.com/Le0nar/find-music/core-service/internal/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// initial grpc connect
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
	 	log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
	}
	defer conn.Close()

	handler := handler.NewHandler(conn)

	router := handler.InitRouter()

	http.ListenAndServe(":3000", router)
}