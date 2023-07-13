package handler

import (
	"OpenAI-chat/internal/utilities"
	"html/template"
	"net/http"
)

type Chat struct {
	Ans []string
	Value string
	Size int
}

var global []string
var value string


func Home(w http.ResponseWriter, r *http.Request) {

	tmplt, _ := template.ParseFiles("./templates/form.page.tmpl")

	event := Chat{
		Ans: global,
		Value: value,
	}

	err := tmplt.Execute(w, event)

	if err != nil {
		return
	}
}

func GetData(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("data")
	mainText := utilities.GetData(text)
	global = append(global, mainText)
	value=text
	http.Redirect(w, r, "/", http.StatusFound)
}
