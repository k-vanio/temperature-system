package main

import (
	"log"

	"github.com/k-vanio/temperature-system/internal/core/search"
	"github.com/k-vanio/temperature-system/internal/domain"
	"github.com/k-vanio/temperature-system/internal/infra/server"
)

func main() {

	var zipCode domain.ZipCode = &search.ZipCodeHttp{}
	var actionZipCode domain.Action = server.NewActionZipCode("/", zipCode)
	var app domain.Application = server.New(":8080", []domain.Action{actionZipCode})

	if err := app.Run(); err != nil {
		log.Println(err)
	}
}
