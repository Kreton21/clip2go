package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/boltdb/bolt"
)

func main() {
	createImageHandler := http.HandlerFunc(createImage)
	http.Handle("/send", createImageHandler)
	http.ListenAndServe(":8080", nil)
}

func createImage(w http.ResponseWriter, request *http.Request) {
	db, err := bolt.Open("clip2.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = request.ParseMultipartForm(100 << 20) // maxMemory 100MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Access the photo key - First Approach
	file, h, err := request.FormFile("photo")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tmpfile, err := os.Create("./" + h.Filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer tmpfile.Close()

	_, err = io.Copy(tmpfile, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)

}
