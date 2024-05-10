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

// user handlers
func (app *application) CreateUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Creating user here"))
} 

func (app *application) AddTweet(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Creating post here"))
}

func (app *application) SeeTweets(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("You should be able see list of tweets here"))
}

func (app *application) DeleteTweet(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Deleting tweet here"))
}

func (app *application) UpdateTweet(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Updating tweet here"))
}

func (app *application) AddComment(w http.ResponseWriter , r *http.Request){
	w.Write([]byte("Adding comment here"))
}

func (app *application) DeleteComment(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Creating user here"))
}