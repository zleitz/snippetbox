package main

import (
	"net/http"

	"github.com/zleitz/snippetbox/cmd/web/config"
)

func serverError(app *config.Application) func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		var (
			method = r.Method
			uri    = r.URL.RequestURI()
		)

		app.Logger.Error(err.Error(), "method", method, "uri", uri)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func clientError(app *config.Application) func(http.ResponseWriter, int) {
	return func(w http.ResponseWriter, status int) {
		http.Error(w, http.StatusText(status), status)
	}
}
