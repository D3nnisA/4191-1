// Filename: cmd/api/routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	//Create a new router

	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodPost, "/v1/courses", app.createCoursesHandler)
	router.HandlerFunc(http.MethodGet, "/v1/courses/:id", app.showCoursesHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/courses/:id", app.updateCoursesHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/courses/:id", app.deleteCoursesHandler)

	// return the router
	return router

}
