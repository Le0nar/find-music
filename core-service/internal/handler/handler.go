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
)


type Handler struct {
	gRPCCon *grpc.ClientConn
}

func NewHandler(conn *grpc.ClientConn) *Handler {
	return &Handler{gRPCCon: conn}
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

	c := musicv1.NewMusicClient(h.gRPCCon)
   
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
   
	grpcRequest, err := c.GetMusic(ctx, &musicv1.GetMusicRequest{SearchValue: searchQuery})
	if err != nil {
	 log.Fatalf("error calling function GetMusic: %v", err)
	}
   
	// TODO: now it returns music id
	render.JSON(w, r, grpcRequest.GetId())
}