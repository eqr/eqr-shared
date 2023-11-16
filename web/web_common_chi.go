package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func ShowErrorMessageChi(w http.ResponseWriter, message string, tmpts *template.Template) {
	if err := tmpts.ExecuteTemplate(w, "error.html", map[string]interface{}{
		"message": message,
	}); err != nil {
		err1 := fmt.Errorf("cannot execute template error.html: %w", err)
		log.Println(err1)
	}
}

func ShowErrorChi(w http.ResponseWriter, err error, tmplts *template.Template) {
	ShowErrorMessageChi(w, err.Error(), tmplts)
}

func RedirectChi(w http.ResponseWriter, r *http.Request, targetUrl string) {
	http.Redirect(w, r, targetUrl, http.StatusFound)
}
