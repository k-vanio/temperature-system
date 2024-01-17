package search

import "github.com/k-vanio/temperature-system/internal/dto"

type ZipCodeHttp struct{}

func (z ZipCodeHttp) Search(request dto.SearchRequest) (dto.SearchResponse, error) {
	return dto.SearchResponse{}, nil
}
