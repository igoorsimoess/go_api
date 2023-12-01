package handlers

import (
	"github.com/go-chi/chi"
	"github.com/igoorsimoess/go_api/internal/middleware"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux){
	// middleware - a function that is called first before the primary function to handle something
	// Global middleware

	r.Use(chimiddle.StripSlashes) // make sure all the trailing slashes are ignored

	r.Route("/account", func(router chi.Router){ // anonymous functions that takes a chi router as second parameter
		// checks if the user is authorized.
		// implemented on //middleware module
		router.Use(middleware.Authorization)
		// throws an error and then the rest of the code won't be executed
		router.Get("/coins", GetCoinBalance)

	})
}
