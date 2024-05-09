package main

import (
	"fmt"
	"log"
	"net/http"
)

const version = "1.0.0"

type config struct{
	port int
	env string
}

type application struct {
	config config
	log log.Logger
}

func main() {
	fmt.Println("Hello Api")

}