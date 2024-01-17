package domain

import "github.com/k-vanio/temperature-system/internal/dto"

type ZipCode interface {
	Search(request dto.SearchRequest) (dto.SearchResponse, error)
}
