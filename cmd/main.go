package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	addr := ":80"
	log.Printf("listening on port%s\n", addr)

	e := echo.New()
	e.Static("/static", "public")
	e.File("/", "public/index.html")

	err := http.ListenAndServe(addr, e)
	if err != nil {
		log.Printf("server got terminated, err: %s\n", err.Error())
	}
}
