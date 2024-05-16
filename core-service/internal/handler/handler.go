package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	musicv1 "github.com/Le0nar/find-music-protos/gen/go/music"
	"github.com/Le0nar/find-music/core-service/internal/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Handler struct {
	musicClient musicv1.MusicClient
}

func NewHandler(musicClient musicv1.MusicClient) *Handler {
	return &Handler{musicClient: musicClient}
}

// Initialization of router
func (h *Handler) InitRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/music", h.GetMusic)
	router.Post("/music", h.CreateMusic)

	return router
}

func (h* Handler) GetMusic(w http.ResponseWriter, r *http.Request)  {
	searchQuery := r.URL.Query().Get("searchQuery")
   
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
   
	grpcRequest, err := h.musicClient.GetMusic(ctx, &musicv1.GetMusicRequest{SearchValue: searchQuery})
	if err != nil {
	 	log.Fatalf("error calling function GetMusic: %v", err)
	}
   
	// TODO: now it returns music id
	render.JSON(w, r, grpcRequest.GetId())
}

func (h *Handler) CreateMusic(w http.ResponseWriter, r *http.Request) {
	var dto models.MusicDto

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	updatedDto := UppercaseSinger(&dto)

	grpcRequest, err := h.musicClient.CreateMusic(ctx, &musicv1.CreateMusicRequest{Singer: updatedDto.Singer, Track: updatedDto.Track})
	if err != nil {
		log.Fatalf("error calling function CreateMusic: %v", err)
   	}

	render.JSON(w, r, grpcRequest.GetId())
}


// example method for test covering
func UppercaseSinger(originDto *models.MusicDto) models.MusicDto {
	fmt.Printf("originDto: %v\n", originDto)

	var updatedMusicDto models.MusicDto

	uppercasedSinger := strings.ToUpper(originDto.Singer)

	updatedMusicDto.Singer = uppercasedSinger
	updatedMusicDto.Track = originDto.Track

	return updatedMusicDto
}