package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/mbolis/yello/db"
	"github.com/mbolis/yello/model"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/lists", func(w http.ResponseWriter, r *http.Request) {
		ls, err := db.GetAllLists()
		if err != nil {
			w.WriteHeader(500)
			return
		}

		render.JSON(w, r, ls)
	})
	r.Post("/lists", func(w http.ResponseWriter, r *http.Request) {
		var l model.List
		if err := render.DecodeJSON(r.Body, &l); err != nil {
			w.WriteHeader(400)
			return
		}

		if err := db.AddList(&l); err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(201)
		render.JSON(w, r, l)
	})
	if err := http.ListenAndServe(":5000", r); err != nil {
		fmt.Printf("Whoops %e\n", err)
	}
}
