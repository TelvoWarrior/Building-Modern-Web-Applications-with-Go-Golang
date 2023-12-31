package handlers

import (
	"net/http"
	"telvo/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "about.html")
}