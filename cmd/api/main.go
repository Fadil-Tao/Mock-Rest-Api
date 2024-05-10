package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct{
	port int
	env string
}

type application struct {
	config config
	log *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port,"port",8080,"API SERVER PORT")
	flag.StringVar(&cfg.env,"environtment","developement", "Environtment (development|staging|production)")
	flag.Parse()

	infolog := log.New(os.Stdout,"INFORMASI :", log.Ldate | log.Ltime)

	app := &application{
		config: cfg,
		log: infolog,
	}

	

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d",cfg.port),
		Handler: app.Router(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	infolog.Printf("Starting %s server on %s",cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	infolog.Fatal(err)


}