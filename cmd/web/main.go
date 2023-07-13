package main

import (
	"OpenAI-chat/internal/handler"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	mux := chi.NewRouter()
	mux.Get("/", handler.Home)
	mux.Post("/",handler.GetData)

	http.ListenAndServe(":8080", mux)
}




