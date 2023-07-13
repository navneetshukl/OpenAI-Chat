package render

import (
	"OpenAI-chat/internal/models"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, url string, td *models.TemplateData) {
	tmpl, _ := template.ParseFiles("./templates/" + url)

	tmpl.Execute(w, td)
}
