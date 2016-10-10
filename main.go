package main

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
)

var repoHelper *RepoHelper

func main() {
	router := NewRouter()

	repoHelper = NewRepoHelper(getOneSession())

	log.Fatal(http.ListenAndServe(":9001", router))
}

func getOneSession() *mgo.Session {
	log.Printf("getOneSession was called.")

	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return s
}
