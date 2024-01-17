package server

import (
	"log"
	"net/http"

	"github.com/k-vanio/temperature-system/internal/domain"
)

type application struct {
	Port    string
	Actions []domain.Action
}

func New(Port string, Actions []domain.Action) *application {
	return &application{
		Port:    Port,
		Actions: Actions,
	}
}

func (a *application) Run() error {

	mux := http.NewServeMux()

	for _, action := range a.Actions {
		mux.HandleFunc(action.Path(), action.Handle)
	}

	srv := &http.Server{Addr: a.Port, Handler: mux}

	log.Printf("Server running on port %s", a.Port)
	return srv.ListenAndServe()
}
