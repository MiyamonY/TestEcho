///
// File:  main.go
// Author: ymiyamoto
//
// Created on Wed Jun  6 22:04:39 2018
//
package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.String(http.StatusOK, fmt.Sprintf("Hello, %s!", name))
	})

	e.GET("/users/:id", func(c echo.Context) error {
		username := c.Param("id")
		jsonMap := map[string]string{
			"name": username,
			"bar":  "barbar",
		}
		return c.JSON(http.StatusOK, jsonMap)
	})

	e.Logger.Fatal(e.Start(":1234"))
}
