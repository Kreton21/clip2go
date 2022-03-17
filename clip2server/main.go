package main

import (
	"clip2serv/httstuff"
	"net/http"
)

func main() {

	createImageHandler := http.HandlerFunc(httstuff.CreateImage)
	createTextHandler := http.HandlerFunc(httstuff.CreateText)
	getUserClip := http.HandlerFunc(getClip)

	http.Handle("/sendimg", createImageHandler)
	http.Handle("/sendtxt", createTextHandler)
	http.Handle("/get", getUserClip)

	http.ListenAndServe(":8080", nil)
}

func getClip(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("Received a GET request\n"))
}
