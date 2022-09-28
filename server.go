package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/subosito/gotenv"
)

type Name struct {
	NConst    string `json:"nconst"`
	Name      string `json:"name"`
	DeathYear string `json:"deathYear"`
	BirthYear string `json:"birthYear"`
}

func init() {
	gotenv.Load()
}

type Error struct {
	Message string `json:"error"`
}

func main() {
	db, err := NewPostgreSQL()
	if err != nil {
		log.Fatalf("Could not initialize Database connection %s", err)
	}

	defer db.Close()

	mc, err := NewMemcached()

	if err != nil {
		log.Fatalf("Could not initialize Memcached client %s", err)
	}

	router := chi.NewRouter()

	renderJSON := func(w http.ResponseWriter, val interface{}, statusCode int) {
		w.WriteHeader(statusCode)
		_ = json.NewEncoder(w).Encode(val)
	}

	router.HandleFunc("/names/{id}", func(w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(r, "id")

		val, err := mc.GetName(id)
		if err == nil {
			renderJSON(w, &val, http.StatusOK)
			return
		}

		name, err := db.FindByNConst(id)
		if err != nil {
			renderJSON(w, &Error{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		_ = mc.SetName(name)
		renderJSON(w, &name, http.StatusOK)

	})

	fmt.Println("starting server :8080")

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
