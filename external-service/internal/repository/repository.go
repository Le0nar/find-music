package repository

import (
	"context"
	"fmt"
	"math/rand"
	"strings"

	"github.com/Le0nar/find-music/external-service/internal/models"
)

type Repository struct {
	musicList []models.Music
}

func New() *Repository {
	return &Repository{musicList: []models.Music{}}
}

func (r *Repository) CreateMusic(ctx context.Context, singer, track string) (int64, error) {
	createdMusic := models.Music{Singer: singer, Track: track, Id: rand.Int63()}
	extendedMusicList := append(r.musicList, createdMusic)
	r.musicList = extendedMusicList

	fmt.Printf("r.musicList: %v\n", r.musicList)
	return createdMusic.Id, nil
}

func (r *Repository) GetMusic(ctx context.Context, searchQuery string) (*models.Music, error) {
	var findedMusic models.Music
	for i := 0; i < len(r.musicList); i++ {
		if strings.HasPrefix(r.musicList[i].Track, searchQuery) {
			findedMusic = r.musicList[i]
			break
		}
	}
	return &findedMusic, nil
}

func (r *Repository) UpdateMusic(ctx context.Context, singer, track string, id int64) (int64, error) {
	updatedMusic := models.Music{Singer: singer, Track: track, Id: id}
	for i := 0; i < len(r.musicList); i++ {
		if updatedMusic.Id == id {
			r.musicList[i] = updatedMusic
			break
		}
	}
	return id, nil
}

func (r *Repository) DeleteMusic(ctx context.Context, id int64) (int64, error) {
	var findedIndex int
	for i := 0; i < len(r.musicList); i++ {
		if r.musicList[i].Id == id {
			findedIndex = i
			break
		}
	}

	remove(r.musicList, findedIndex)
	return id, nil
}

func remove(s []models.Music, i int) []models.Music {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}