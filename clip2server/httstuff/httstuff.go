package httstuff

import (
	"clip2serv/bolted"
	"fmt"
	"net/http"
)

func CreateText(w http.ResponseWriter, request *http.Request) {
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
func CreateImage(w http.ResponseWriter, request *http.Request) {
	bucket := "bucket"
	key := "image"
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
