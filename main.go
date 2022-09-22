package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/vfreex/art-api/docs/art_api"
	"net/http"
)

// @title Example ART API Server
// @version 1.0
// @description This is a sample ART API server.
// @termsOfService http://swagger.io/terms/

// @contact.name OpenShift Team ART
// @contact.url https://example.com/
// @contact.email aos-team-art@redhat.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/releases/:name", getRelease)
	e.Logger.Fatal(e.Start(":8080"))
}

type Release struct {
	Name     string   `json:"name" xml:"name" form:"name" query:"name"`
	Accepted bool     `json:"accepted" xml:"accepted" form:"accepted" query:"accepted"`
	Assembly Assembly `json:"assembly" xml:"assembly" form:"assembly" query:"assembly"`
}
type Assembly struct {
	Type string `json:"type" xml:"type" form:"type" query:"type"`
}

// @Summary Get release info
// @Description Get the info of specified release
// @Tags root
// @Accept */*
// @Produce json
// @Param id path string  true  "Release name"
// @Success 200 {object} Release
// @Router /releases/{name} [get]
func getRelease(c echo.Context) error {
	name := c.Param("name")
	if name == "4.10.33" {
		release := Release{
			Name:     "4.10.33",
			Accepted: true,
			Assembly: Assembly{Type: "standard"},
		}
		return c.JSON(http.StatusOK, release)
	}
	return c.String(http.StatusNotFound, "Release not found.")
}
