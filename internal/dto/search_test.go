package dto_test

import (
	"testing"

	"github.com/k-vanio/temperature-system/internal/dto"
)

func TestSearch(t *testing.T) {
	t.Run("Test SearchRequest", func(t *testing.T) {
		request := dto.SearchRequest{ZipCode: "07987110"}
		if request.ZipCode != "07987110" {
			t.Errorf("Expected ZipCode to be 07987110, got %s", request.ZipCode)
		}
	})

	t.Run("Test SearchResponse", func(t *testing.T) {
		response := dto.SearchResponse{Status: 200, Body: "Test Body"}

		if response.Status != 200 {
			t.Errorf("Expected Status to be 200, got %v", response.Status)
		}

		if response.Body != "Test Body" {
			t.Errorf("Expected Body to be Test Body, got %v", response.Body)
		}
	})
}
