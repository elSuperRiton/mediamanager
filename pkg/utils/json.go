package utils

import (
	"net/http"

	"github.com/go-chi/render"
)

func setContentTypeJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// FormatErr return a formated json response with
// err and statusCode attached
func FormatErr(msg string, code int) interface{} {
	return map[string]interface{}{
		"status": code,
		"error":  msg,
	}
}

// RenderErr uses go-chi render package in order to render
// a properly formated error message
func RenderErr(w http.ResponseWriter, r *http.Request, msg string, code int) {
	setContentTypeJSON(w)
	w.WriteHeader(code)
	render.JSON(w, r, FormatErr(msg, code))
	return
}

// FormatData returns a formated json response with data
// and statusCode attached
func FormatData(data interface{}, code int) interface{} {
	return map[string]interface{}{
		"status": code,
		"data":   data,
	}
}

// RenderData uses go-chi render package in order to render
// a properly formatted success message
func RenderData(w http.ResponseWriter, r *http.Request, data interface{}, code int) {
	setContentTypeJSON(w)
	w.WriteHeader(code)
	render.JSON(w, r, FormatData(data, code))
}
