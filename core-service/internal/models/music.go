package models

type Music struct {
	Singer string
	Track  string
	Id     int64
}

type MusicDto struct {
	Singer string `json:"singer" binding:"required"`
	Track  string `json:"track" binding:"required"`
}