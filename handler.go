package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Todo struct {
	Name      string    `json:"name`
	Completed bool      `json:completed`
	Due       time.Time `json:due`
}
type Todos []Todo

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoItems(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}
