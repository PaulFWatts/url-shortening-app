package controllers

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/PaulFWatts/url-shortening-app/url"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL can't be empty", http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(originalURL, "http") && !strings.HasPrefix(originalURL, "https") {
		originalURL = "https://" + originalURL
	}

	shortURL := url.Shorten(originalURL)

	data := map[string]string{
		"ShortURL": shortURL,
	}

	t, err := template.ParseFiles("internal/views/shorten.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
