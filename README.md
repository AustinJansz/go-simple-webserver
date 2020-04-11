# go-simple-webserver

*Author: Austin Jansz - DURHAM COLLEGE | UOIT | OntarioTechU*

An extensible Echo webserver with bad-request logging.

- Intended Purpose:
    - Security investigation aid that offers simple logging
    - Extensible with [Middlewares](https://echo.labstack.com/middleware)

- Notes:
    - Capable of hosting static sites (see sample site)
    - Default path setting : current directory
    - Default port setting : 8080
    - Logging is enabled for failed requests

## Instructions

### Simple Method

0. Place the executable in a directory with the files that are to be hosted
1. Run the executable binary for the desired system
2. Access port 8080 via a browser

### Customized Operation

0. ./go-simple-webserver(.exe) -path *path to directory* -port *port number*