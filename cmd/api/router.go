package main

import "net/http"

func (app *application) Router() http.Handler{
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", app.HealthCheckHandler)
	mux.HandleFunc("GET /tweet/list" , app.SeeTweets)
	mux.HandleFunc("POST /tweet/create", app.AddTweet)
	mux.HandleFunc("PUT /tweet/edit/", app.UpdateTweet)
	mux.HandleFunc("DELETE /tweet/delete/", app.DeleteTweet)
	mux.HandleFunc("POST /tweet/comment/create", app.AddComment)

	v1 := http.NewServeMux()
	v1.Handle("/v1/" , http.StripPrefix("/v1", mux))

	return v1	
}