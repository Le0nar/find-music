package service

import (
	"context"

	"github.com/Le0nar/find-music/external-service/internal/models"
)

type repository interface {
	CreateMusic(ctx context.Context, singer, track string) (int64, error)
	GetMusic(ctx context.Context, searchQuery string) (*models.Music, error)
	UpdateMusic(ctx context.Context, singer, track string, id int64) (int64, error)
	DeleteMusic(ctx context.Context, id int64) (int64, error)
}

type Service struct {
	repository repository
}

func New(repo repository) *Service {
	return &Service{repository: repo}
}

func (s *Service) CreateMusic(ctx context.Context, singer, track string) (int64, error) {
	return s.repository.CreateMusic(ctx, singer, track)
}

func (s *Service) GetMusic(ctx context.Context, searchQuery string) (*models.Music, error) {
	return s.repository.GetMusic(ctx, searchQuery)
}

func (s *Service) UpdateMusic(ctx context.Context, singer, track string, id int64) (int64, error) {
	return s.repository.UpdateMusic(ctx, singer, track, id)
}

func (s *Service) DeleteMusic(ctx context.Context, id int64) (int64, error) {
	return s.repository.DeleteMusic(ctx, id)
}
