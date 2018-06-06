///
// File:  main.go
// Author: ymiyamoto
//
// Created on Wed Jun  6 22:04:39 2018
//
package main

import (
	"github.com/MiyamonY/TestEcho/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler.Index)
	e.GET("/users/:id", handler.Users)
	e.GET("/simplehtml", handler.SimpleHTML)

	e.Logger.Fatal(e.Start(":1234"))
}
