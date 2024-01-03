package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.notFoundResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/wines", app.listWinesHandler)
	router.HandlerFunc(http.MethodGet, "/v1/wines/:id", app.getWineHandler)
	router.HandlerFunc(http.MethodPost, "/v1/wines", app.createWineHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/wines/:id", app.updateWineHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/wines/:id", app.deleteWineHandler)

	return router
}
