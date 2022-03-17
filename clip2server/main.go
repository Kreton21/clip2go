package main

import (
	"clip2serv/bolted"
	"fmt"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("clip2.db", 0600, nil)
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("bucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
	createImageHandler := http.HandlerFunc(createImage)
	createTextHandler := http.HandlerFunc(createText)

	getUserClip := http.HandlerFunc(getClip)
	http.Handle("/sendimg", createImageHandler)
	http.Handle("/sendtxt", createTextHandler)

	http.Handle("/get", getUserClip)
	http.ListenAndServe(":8080", nil)
}
func getClip(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("Received a GET request\n"))
}

func createText(w http.ResponseWriter, request *http.Request) {
	bucket := "bucket"
	key := "image"
	//var value = "123test"
	err := request.ParseMultipartForm(100 << 20) // maxMemory 100MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Access the photo key - First Approach
	value := request.FormValue("photo")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bolted.Wdb([]byte(bucket), []byte(key), []byte(value))
	var a = bolted.Rdb(bucket, key)

	fmt.Print(a)

	w.WriteHeader(200)
}
func createImage(w http.ResponseWriter, request *http.Request) {
	bucket := "bucket"
	key := "image"
	//var value = "123test"
	err := request.ParseMultipartForm(100 << 20) // maxMemory 100MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Access the photo key - First Approach
	value := request.FormValue("photo")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bolted.Wdb([]byte(bucket), []byte(key), []byte(value))
	var a = bolted.Rdb(bucket, key)

	fmt.Println(a)

	w.WriteHeader(200)

}
