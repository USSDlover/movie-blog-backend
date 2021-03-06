package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovies)

	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.getOneMovie)

	return app.enableCORS(router)
}
