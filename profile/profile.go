package profile

import (
	"log"
	"net/http"
)

type PlayerProfile struct {
	username string
	rank     string
	agents   []string
}

func ProfilePageHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	profiles := []PlayerProfile{{
		"vectorini",
		"P1",
		[]string{"Omen", "Cypher"},
	}}
	component := Page(profiles)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in ProfileListHandler: %e", err)
	}
}
