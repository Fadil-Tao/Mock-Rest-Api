package main

import (
	"fmt"
	"net/http"
)

func (app *application) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: Tersedia")
	fmt.Fprintf(w, "environment: %s\n ", app.config.env )
	fmt.Fprintf(w, "Version : %s\n", version)
}


