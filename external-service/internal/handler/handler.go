package handler

import (
	"context"

	musicv1 "github.com/Le0nar/find-music-protos/gen/go/music"
	"github.com/Le0nar/find-music/external-service/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const emptyValue = 0

type musicService interface {
	CreateMusic(ctx context.Context, singer, track string) (int64, error)
	GetMusic(ctx context.Context, searchQuery string) (*models.Music, error)
	UpdateMusic(ctx context.Context, singer, track string, id int64) (int64, error)
	DeleteMusic(ctx context.Context, id int64) (int64, error)
}

type Handler struct {
	musicv1.UnimplementedMusicServer
	MusicService musicService
}

func New (service musicService) *Handler {
	return &Handler{MusicService: service}
}

func (h *Handler) CreateMusic(ctx context.Context, in *musicv1.CreateMusicRequest) (*musicv1.CreateMusicResponse, error) {
	if in.Singer == "" {
		return nil, status.Error(codes.InvalidArgument, "singer is required")
	}
	if in.Track == "" {
		return nil, status.Error(codes.InvalidArgument, "track is required")
	}

	// TODO: i dont know what is better: "in.Singer" or "in.GetSinger()"
	id, err := h.MusicService.CreateMusic(ctx, in.GetSinger(), in.GetTrack())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create the music")
	}

	return &musicv1.CreateMusicResponse{Id: id}, nil
}

func (h *Handler) GetMusic(ctx context.Context, in *musicv1.GetMusicRequest) (*musicv1.GetMusicResponse, error) {
	if in.SearchValue == "" {
		return nil, status.Error(codes.InvalidArgument, "search value is required")
	}

	// TODO: rename
	musicInstance, err := h.MusicService.GetMusic(ctx, in.GetSearchValue())
	if err != nil {
		// TODO: mb handle 2 cases: internal error and music not found
		return nil, status.Error(codes.Internal, "internal Handler error")
	}

	return &musicv1.GetMusicResponse{Singer: musicInstance.Singer, Track: musicInstance.Track, Id: musicInstance.Id}, nil
}

func (h *Handler) UpdateMusic(ctx context.Context, in *musicv1.UpdateMusicRequest) (*musicv1.UpdateMusicResponse, error) {
	if in.Singer == "" {
		return nil, status.Error(codes.InvalidArgument, "singer is required")
	}
	if in.Track == "" {
		return nil, status.Error(codes.InvalidArgument, "track is required")
	}
	if in.Id == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	id, err := h.MusicService.UpdateMusic(ctx, in.GetSinger(), in.GetTrack(), in.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to update the music")
	}

	return &musicv1.UpdateMusicResponse{Id: id}, nil
}

func (h *Handler) DeleteMusic(ctx context.Context, in *musicv1.DeleteMusicRequest) (*musicv1.DeleteMusicResponse, error) {
	if in.Id == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	id, err := h.MusicService.DeleteMusic(ctx, in.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete the music")
	}

	return &musicv1.DeleteMusicResponse{Id: id}, nil
}