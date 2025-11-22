package handlers

import (
	"html/template"
	"net/http"

	"ascii-art-web/internal/service"
)

// PostHandler обрабатывает и GET-запросы к "/", и POST-запросы к "/ascii-art"
func PostHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, "Template parse error")
		return
	}

	switch r.URL.Path {
	case "/":
		if r.Method != http.MethodGet {
			ErrorPage(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		data := PageData{
			Title: "ASCII Text Generator",
		}

		err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			ErrorPage(w, http.StatusInternalServerError, "Template execution error")
			return
		}

	case "/ascii-art":
		if r.Method != http.MethodPost {
			ErrorPage(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		if err := r.ParseForm(); err != nil {
			IndexError(w, tmpl, "Failed to parse form")
			return
		}

		text := r.FormValue("text")
		banner := r.FormValue("banner")

		if len(text) == 0 || len(banner) == 0 {
			IndexError(w, tmpl, "Failed to generate ASCII art")
			return
		}

		dicitonary, err := service.LoadBanner(banner)
		if err != nil {
			IndexError(w, tmpl, err.Error())
			return
		}

		result, err := service.PrintASCII(text, dicitonary)
		if err != nil {
			IndexError(w, tmpl, err.Error())
			return
		}

		data := PageData{
			Title:       "ASCII Text Generator",
			Banner:      banner,
			InputText:   text,
			AsciiResult: result,
		}

		w.WriteHeader(http.StatusOK)
		err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			IndexError(w, tmpl, err.Error())
			return
		}

	default:
		ErrorPage(w, http.StatusNotFound, "Page not found")
	}
}

func ErrorPage(w http.ResponseWriter, status int, msg string) {
	tmplErr, _ := template.ParseFiles("templates/error.html")
	info := ErrInfo{
		Title:     "Error-Page",
		ErrorType: status,
		ErrorMsg:  msg,
	}

	w.WriteHeader(status)
	tmplErr.ExecuteTemplate(w, "error", info)
}

func IndexError(w http.ResponseWriter, tmpl *template.Template, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	data := PageData{
		Title:    "ASCII Text Generator",
		ErrorMsg: msg,
	}

	tmpl.ExecuteTemplate(w, "index", data)
}
