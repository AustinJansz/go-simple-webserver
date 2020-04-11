package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// global variables
var (
	path string = ""
	port string = ""
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
	e.Use(middleware.Static(path))
	// Custom middleware that sets the headers of the server
	e.Use(serverHeader)
	// Interaction logger that allows for the checking of requests
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}]   ${status}   ${method}   ${host}${path}\n",
	}))
}

// flagHandler : capture flags from console and set their helper messages
func flagHandler() {
	// Capture the user flags path and port
	flag.StringVar(&path, "path", "./", "Directory of files to serve")
	flag.StringVar(&port, "port", "8080", "Server listen port")
	flag.Parse()
}

// main : main entrypoint for the progarm
func main() {
	flagHandler()
	// Check to see that port is an integer
	if _, err := strconv.Atoi(port); err != nil {
		fmt.Println("Number required for port")
		return
	}
	// Format the port as an address
	port = ":" + port

	fmt.Println("Welcome to the server, we live to serve.")
	fmt.Println("The site's entrypoint path is " + path)
	// Create the new server
	e := echo.New()
	// Install the middlewares into the echo instance
	setMiddlewares(e)
	e.Start(port)
}
