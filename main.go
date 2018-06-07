///
// File:  main.go
// Author: ymiyamoto
//
// Created on Wed Jun  6 22:04:39 2018
//
package main

import (
	"database/sql"
	"html/template"
	"io"

	"github.com/MiyamonY/TestEcho/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

type templates struct {
	templates *template.Template
}

func (t *templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		e.Logger.Fatal("sql error:", err)
	}
	defer db.Close()

	h := handler.NewHandler(db)

	e.Static("/assets", "assets")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := &templates{templates: template.Must(template.ParseGlob("assets/views/*.html"))}
	e.Renderer = t

	e.GET("/", h.Index)
	e.GET("/users/:id", h.Users)
	e.GET("/simplehtml", h.SimpleHTML)
	e.GET("/html", h.Template)
	e.POST("/html", h.Post)

	e.Logger.Fatal(e.Start(":1234"))
}
