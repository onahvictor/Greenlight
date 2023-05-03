package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	//initialiaing a our httprouter
	mux := httprouter.New()
	mux.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFoundResponse(w, r)
	})
	mux.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.methodNotAllowedResponse(w, r)
	})

	mux.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	mux.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	mux.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	return mux
}
