///
// File:  handler.go
// Author: ymiyamoto
//
// Created on Wed Jun  6 22:46:14 2018
//
package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Name struct {
	First string `form:"first"`
	Last  string `form:"last"`
}

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) Handler {
	return Handler{db: db}
}

func (h *Handler) Index(c echo.Context) error {
	name := c.QueryParam("name")
	return c.String(http.StatusOK, fmt.Sprintf("Hello, %s!", name))
}

func (h *Handler) Users(c echo.Context) error {
	username := c.Param("id")
	jsonMap := map[string]string{
		"name": username,
		"bar":  "barbar",
	}
	return c.JSON(http.StatusOK, jsonMap)
}

func (h *Handler) SimpleHTML(c echo.Context) error {
	return c.HTML(http.StatusOK, "<strong>Hello, Wold!</strong>")
}

func (h *Handler) Template(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "world")
}

func (h *Handler) Post(c echo.Context) error {
	n := &Name{} // important: must be pointer
	if err := c.Bind(n); err != nil {
		return err
	}

	if _, err := h.db.Exec("INSERT INTO users (last, first) VALUES (?, ?)", n.First, n.Last); err != nil {
		return c.HTML(http.StatusInternalServerError, "Ooops! Server Error")
	}

	return c.Redirect(http.StatusMovedPermanently, "html")
}
