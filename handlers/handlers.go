package handlers

import (
	a "ascii-art-web/ascii-art/ascii-art"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

type Result struct {
	Res  string
	Post bool
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/ascii-art" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	for _, v := range text {
		if (v >= ' ' && v <= '~') || v == '\r' || v == '\n' {
			continue
		} else {
			ErrorHandler(w, http.StatusBadRequest)
			return
		}
	}
	if len(text) > 300 {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	if text == "" || banner == "" {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	res, status := a.MainAsciiArt(text, banner)
	if status != 200 && StatusError(w, status) {
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	res1 := Result{
		Res:  res,
		Post: true,
	}
	err = tmpl.Execute(w, res1)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func StatusError(w http.ResponseWriter, status int) bool {
	if status == http.StatusNotFound {
		ErrorHandler(w, http.StatusNotFound)
		return true
	}
	if status == http.StatusInternalServerError {
		ErrorHandler(w, http.StatusInternalServerError)
		return true
	}
	if status == http.StatusBadRequest {
		ErrorHandler(w, http.StatusBadRequest)
		return true
	}
	return false
}

func ErrorHandler(w http.ResponseWriter, status int) {
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	ErrorMessage := ErrorStatus{status, http.StatusText(status)}
	w.WriteHeader(status)
	err = tmpl.Execute(w, ErrorMessage)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

type ErrorStatus struct {
	Code    int
	Message string
}
