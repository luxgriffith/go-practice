package adventure_view

import (
	"html/template"
	"net/http"
	"os"
)

// Driver function, takes in an error object and builds up the error page
func BuildErrorPage(w http.ResponseWriter, err error) error {
	if writeErr := createTemplate(err); writeErr != nil {
		return writeErr
	}
	templ, _ := template.ParseFiles("./error-page-template.html", "./error-page.hml")
	templ.ExecuteTemplate(w, "layout", nil)
	return nil
}

// Creates the template file that we'll use to template the error page
func createTemplate(err error) error {
	msg := err.Error()
	templateStr := "{{define \"message\"}} " + msg + "{{end}}"
	return os.WriteFile("error-page-template.html", []byte(templateStr), 0644)
}
