package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/igoorsimoess/go_api/internal/handlers"
	log "github.com/sirupsen/logrus" // aliased as log
	// go mod tidy
)

func main(){
	log.SetReportCaller(true) // when we print we get the file and the line number

	var r *chi.Mux = chi.NewRouter() // returns a pointer to a mux type.
	// It's a struct we'll use to setup the api
	handlers.Handler(r)

	fmt.Println("Starting GO api")

	// catching potential api error
	err := http.ListenAndServe("localhost:8000", r)

	if err != nil{
		log.Error(err)
	}
}