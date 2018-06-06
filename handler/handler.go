///
// File:  handler.go
// Author: ymiyamoto
//
// Created on Wed Jun  6 22:46:14 2018
//
package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	name := c.QueryParam("name")
	return c.String(http.StatusOK, fmt.Sprintf("Hello, %s!", name))
}

func Users(c echo.Context) error {
	username := c.Param("id")
	jsonMap := map[string]string{
		"name": username,
		"bar":  "barbar",
	}
	return c.JSON(http.StatusOK, jsonMap)
}

func SimpleHTML(c echo.Context) error {
	return c.HTML(http.StatusOK, "<strong>Hello, Wold!</strong>")
}

func Template(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "world")
}
