package profile

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PlayerProfile struct {
	username string
	rank     string
	agents   []string
}

func ProfileListHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	profiles := []PlayerProfile{{
		"vector#ini",
		"P1",
		[]string{"Omen", "Cypher"},
	}}
	component := ProfilesPage(profiles)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in ProfileListHandler: %e", err)
	}
}

func ProfilePageHandler(c echo.Context) error {
	username := c.Param("username")
	// todo: load profile from db
	profiles := []PlayerProfile{{
		"vector#ini",
		"P1",
		[]string{"Omen", "Cypher"},
	}}
	for _, profile := range profiles {
		if profile.username == username {
			component := ProfileOverviewPage(profile)
			return component.Render(c.Request().Context(), c.Response().Writer)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound, "Profile not found")
}
