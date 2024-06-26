package handler

import (
	"testing"

	"github.com/Le0nar/find-music/core-service/internal/models"
	"github.com/stretchr/testify/assert"
)

// Tesing pure function example
func TestUppercaseSinger(t *testing.T) {
	// Arrange

	testTable := []struct{
		example models.MusicDto
		expected models.MusicDto
	} {
		{
			example: models.MusicDto{Singer: "singer", Track: "track"},
			expected:  models.MusicDto{Singer: "SINGER", Track: "track"},
		},
		{
			example: models.MusicDto{Singer: "keeeek", Track: "track2"},
			expected:  models.MusicDto{Singer: "KEEEEK", Track: "track2"},
		},
		{
			example: models.MusicDto{Singer: "0aaaa1", Track: "track3"},
			expected:  models.MusicDto{Singer: "0AAAA1", Track: "track3"},
		},
	}

	// Act

	for _, item := range testTable {
		result := UppercaseSinger(&item.example)

		// Assert
		assert.Equal(t, item.expected, result)
	}
}
