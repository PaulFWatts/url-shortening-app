package controllers

import (
	"database/sql"
	"html/template"
	"net/http"
	"strings"

	"github.com/PaulFWatts/url-shortening-app/db"
	"github.com/PaulFWatts/url-shortening-app/url"
)

func Shorten(lite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		if err := db.StoreURL(lite, shortURL, originalURL); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

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
}

func Proxy(lite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortUrl := r.URL.Path[1:]
		if shortUrl == "" {
			http.Error(w, "URL can't be empty", http.StatusBadRequest)
			return
		}
		origUrl, err := db.GetOriginalURL(lite, shortUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		http.Redirect(w, r, origUrl, http.StatusPermanentRedirect)
	}
}
