package main

import (
	"log"
	"net/http"
	"CRUD-nethttp/model"
	"CRUD-nethttp/handler"

)

func main() {
	port := 8000
	h := handler.HTTPHandler{Candidates: []*model.Candidate{}}
	http.HandleFunc("/candidates", h.HandleListAndCreate)
	http.HandleFunc("/candidates/", h.HandleDetailAndModify)
	log.Printf("Server currently running at port:%v\n", port)
	log.Fatal(http.ListenAndServe(":8000", nil))
}