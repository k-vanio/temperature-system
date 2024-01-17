package dto

type SearchRequest struct {
	ZipCode string `json:"zip_code"`
}

type SearchResponse struct {
	Status string `json:"status"`
	Body   struct {
		TempC float64 `json:"temp_c"`
		TempF float64 `json:"temp_f"`
		TempK float64 `json:"temp_k"`
	} `json:"body"`
}
