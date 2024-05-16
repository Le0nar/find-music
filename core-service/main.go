package main

import (
	"log"
	"net/http"

	musicv1 "github.com/Le0nar/find-music-protos/gen/go/music"
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
	musicClient := musicv1.NewMusicClient(conn)

	hndlr := handler.NewHandler(musicClient)

	router := hndlr.InitRouter()

	http.ListenAndServe(":3000", router)
}