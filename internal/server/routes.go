package server

import (
	"net/http"

	"valik/profile"
	"valik/web"
	"valik/web/pages"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/assets/*", echo.WrapHandler(fileServer))

	e.GET("/", echo.WrapHandler(templ.Handler(pages.Home())))
	e.GET("/amy", echo.WrapHandler(templ.Handler(pages.Amy())))

	e.GET("/profiles", echo.WrapHandler(http.HandlerFunc(profile.ProfileListHandler)))
	e.GET("/profile/:username", profile.ProfilePageHandler)

	e.GET("/team", echo.WrapHandler(templ.Handler(pages.Team())))

	e.GET("/health", s.healthHandler)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
