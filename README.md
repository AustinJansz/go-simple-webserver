# go-simple-webserver

*Author: Austin Jansz - DURHAM COLLEGE | UOIT | OntarioTechU*

A simple webserver built in Go using LabStack's Echo!

- Intended Purpose:
    - Security investigation aid

- Notes:
    - Capable of hosting static sites (see sample site)
    - Opens on port 8080
    - Logging is enabled for failed requests

- Planned updates:
    - Add help feature (-h/--help)
    - Add arguments (site path, port)

## Instructions

### Simple Method

0. Place the executable in a directory with the files that are to be hosted
1. Run the executable binary for the desired system
2. Access port 8080

### Run from Source

0. Place the source code (main.go) in a directory with the files that are to be hosted
1. go run main.go
