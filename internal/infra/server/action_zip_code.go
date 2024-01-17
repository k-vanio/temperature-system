package server

import (
	"net/http"

	"github.com/k-vanio/temperature-system/internal/domain"
)

type actionZipCode struct {
	path    string
	zipCode domain.ZipCode
}

func NewActionZipCode(path string, zipCode domain.ZipCode) *actionZipCode {
	return &actionZipCode{
		path:    path,
		zipCode: zipCode,
	}
}

func (a actionZipCode) Path() string {
	return a.path
}

func (a actionZipCode) Handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
