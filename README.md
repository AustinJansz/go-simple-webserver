# go-simple-webserver

*Author: Austin Jansz - DURHAM COLLEGE | UOIT | OntarioTechU*

A simple webserver built in Go using LabStack's Echo!

- Intended Purpose:
    - Security investigation aid that offers simple logging
    - Extensible with Middlewares

- Notes:
    - Capable of hosting static sites (see sample site)
    - Opens on port 8080
    - Logging is enabled for failed requests

## Instructions

### Simple Method

0. Place the executable in a directory with the files that are to be hosted
1. Run the executable binary for the desired system
2. Access port 8080 via a browser

### Customized Operation

0. ./go-simple-webserver(.exe) -path *path to director* -port *port number*