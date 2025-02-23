package main

import (
	"net/http"
	"strconv"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strconv.Itoa(200)))
}
