package main

import (
	"net/http"

	"github.com/nullsploit01/spotui-server/internal/response"
)

func (app *application) status(w http.ResponseWriter, r *http.Request) {

	err := response.JSON(w, http.StatusOK, "All Good!")
	if err != nil {
		app.ServerError(w, r, err)
	}
}

func (app *application) sup(w http.ResponseWriter, r *http.Request) {
	err := response.JSON(w, http.StatusOK, "sup?")
	if err != nil {
		app.ServerError(w, r, err)
	}
}
