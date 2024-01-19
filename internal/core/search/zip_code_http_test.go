package search_test

import (
	"net/http"
	"testing"

	"github.com/k-vanio/temperature-system/internal/config"
	"github.com/k-vanio/temperature-system/internal/core/search"
	"github.com/k-vanio/temperature-system/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type zipCodeHttpStub struct {
	mock.Mock
}

func (z *zipCodeHttpStub) Do(request *http.Request) (*http.Response, error) {
	args := z.Called(request)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestZipCodeHttp(t *testing.T) {
	t.Run("should return error when zipCode is invalid", func(t *testing.T) {
		// Arrange
		clientStub := &zipCodeHttpStub{}
		config := config.New()

		zcr := search.New(clientStub, config)

		// Act
		req := dto.SearchRequest{ZipCode: "123456-789"}
		response := zcr.Search(req)

		// Assert
		assert.Equal(t, http.StatusUnprocessableEntity, response.Status)
		assert.EqualValues(t, struct{ Message string }{Message: "invalid zipCode"}, response.Body)
	})
}
