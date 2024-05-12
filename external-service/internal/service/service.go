package service

import (
	"context"

	"github.com/Le0nar/find-music/external-service/internal/models"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) CreateMusic(ctx context.Context, singer, track string) (int64, error) {
	return 0, nil
}

func (s *Service) GetMusic(ctx context.Context, searchQuery string) (*models.Music, error) {
	return &models.Music{}, nil
}

func (s *Service) UpdateMusic(ctx context.Context, singer, track string, id int64) (int64, error) {
	return 0, nil
}

func (s *Service) DeleteMusic(ctx context.Context, id int64) (int64, error) {
	return 0, nil
}