package handler

import (
	"testing"

	"github.com/Le0nar/find-music/core-service/internal/models"
	"github.com/google/go-cmp/cmp"
)

// Tesing pure function example
func TestUppercaseSinger(t *testing.T) {
	// Arrange
	exampleDto := models.MusicDto{Singer: "singer", Track: "track"}
	expected := models.MusicDto{Singer: "SINGER", Track: "track"}

	// Act
	result := UppercaseSinger(&exampleDto)

	// Assert
	isEqueal := cmp.Equal(result, expected)
	if !isEqueal {
		t.Error("Incorrect result")
	}
}