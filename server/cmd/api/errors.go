package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/nullsploit01/spotui-server/internal/response"
	"github.com/nullsploit01/spotui-server/internal/validator"
)

func (app *application) ReportServerError(r *http.Request, err error) {
	var (
		message = err.Error()
		method  = r.Method
		url     = r.URL.String()
		trace   = string(debug.Stack())
	)

	requestAttrs := slog.Group("request", "method", method, "url", url)
	app.logger.Error(message, requestAttrs, "trace", trace)
}

func (app *application) ErrorMessage(w http.ResponseWriter, r *http.Request, status int, message string, headers http.Header) {
	message = strings.ToUpper(message[:1]) + message[1:]

	responseBody := response.ResponseBody{
		Error: true,
		Data:  message,
	}

	err := response.JSONWithHeaders(w, status, responseBody, headers)
	if err != nil {
		app.ReportServerError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) ServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.ReportServerError(r, err)

	message := "The server encountered a problem and could not process your request"
	app.ErrorMessage(w, r, http.StatusInternalServerError, message, nil)
}

func (app *application) NotFound(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	app.ErrorMessage(w, r, http.StatusNotFound, message, nil)
}

func (app *application) MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	app.ErrorMessage(w, r, http.StatusMethodNotAllowed, message, nil)
}

func (app *application) BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	app.ErrorMessage(w, r, http.StatusBadRequest, err.Error(), nil)
}

func (app *application) FailedValidation(w http.ResponseWriter, r *http.Request, v validator.Validator) {
	err := response.JSON(w, http.StatusUnprocessableEntity, v)
	if err != nil {
		app.ServerError(w, r, err)
	}
}

func (app *application) BasicAuthenticationRequired(w http.ResponseWriter, r *http.Request) {
	headers := make(http.Header)
	headers.Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)

	message := "You must be authenticated to access this resource"
	app.ErrorMessage(w, r, http.StatusUnauthorized, message, headers)
}
