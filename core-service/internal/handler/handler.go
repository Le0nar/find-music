package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	musicv1 "github.com/Le0nar/find-music-protos/gen/go/music"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


type Handler struct {}

func NewHandler() *Handler {
	return &Handler{}
}

// Initialization of router
func (h *Handler) InitRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/music", h.GetMusic)

	return router
}

func (h* Handler) GetMusic(w http.ResponseWriter, r *http.Request)  {
	searchQuery := r.URL.Query().Get("searchQuery")

	// TODO: OPEN GRPC CONNECT  ONE TIME IN app.go / main.go file
	
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
	 log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
	}
	defer conn.Close()
	c := musicv1.NewMusicClient(conn)
   
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
   
	grpcRequest, err := c.GetMusic(ctx, &musicv1.GetMusicRequest{SearchValue: searchQuery})
	if err != nil {
	 log.Fatalf("error calling function GetMusic: %v", err)
	}
   
	// TODO: now return music id
	render.JSON(w, r, grpcRequest.GetId())
}