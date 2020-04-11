package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// serverHeader : Custom middleware to change the sever respons header
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Built in header (Server)
		c.Response().Header().Set(echo.HeaderServer, "SGWS/1.0")
		// Custom header
		c.Response().Header().Set("Owner", "nobody")
		return next(c)
	}
}

// setMiddlewares : Control what middlewares are in use by the server
func setMiddlewares(e *echo.Echo) {
	// Points to the path where the server should look for files
	e.Use(middleware.Static("./"))
	// Custom middleware that sets the headers of the server
	e.Use(serverHeader)
	// Interaction logger that allows for the checking of requests
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}]   ${status}   ${method}   ${host}${path}\n",
	}))
}

// main : main entrypoint for the progarm
func main() {
	fmt.Println("Welcome to the server, we live to serve.")
	fmt.Println("The site's entrypoint path is the current directory.")
	// Create the new server
	e := echo.New()
	// Install the middlewares into the echo instance
	setMiddlewares(e)
	e.Start(":8080")
}
