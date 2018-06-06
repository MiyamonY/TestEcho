///
// File:  main.go
// Author: ymiyamoto
//
// Created on Wed Jun  6 22:04:39 2018
//
package main

import (
	"html/template"
	"io"

	"github.com/MiyamonY/TestEcho/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type templates struct {
	templates *template.Template
}

func (t *templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	e.Static("/assets", "assets")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := &templates{templates: template.Must(template.ParseGlob("assets/views/*.html"))}
	e.Renderer = t

	e.GET("/", handler.Index)
	e.GET("/users/:id", handler.Users)
	e.GET("/simplehtml", handler.SimpleHTML)
	e.GET("/html", handler.Template)

	e.Logger.Fatal(e.Start(":1234"))
}
