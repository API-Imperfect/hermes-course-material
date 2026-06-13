package handlers

import (
	"encoding/json"
	"net/http"
)

// RegisterUsers wires the users routes onto the mux.
// Convention: each resource file exposes a Register<Resource>(mux) func
// that main.go calls. Go 1.22 method+path patterns mean no router library.
func RegisterUsers(mux *http.ServeMux) {
	mux.HandleFunc("GET /users", listUsers)
}

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// In-memory only — no database (see AGENTS.md).
var users = []user{
	{ID: 1, Name: "Ada"},
	{ID: 2, Name: "Alan"},
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
