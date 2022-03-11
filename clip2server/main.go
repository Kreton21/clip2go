package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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
	getUserClip := http.HandlerFunc(getClip)
	http.Handle("/send", createImageHandler)
	http.Handle("/get", getUserClip)
	http.ListenAndServe(":8080", nil)
}
func getClip(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("Received a GET request\n"))
}
func Wdb(bucket, key, value []byte) {
	db, err := bolt.Open("clip2.db", 0600, nil)
	if err != nil {
		fmt.Print(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		err := b.Put(key, value)
		return err
	})
	db.Close()
}

func Rdb(bucket, key string) string {
	var result string
	db, err := bolt.Open("clip2.db", 0600, nil)
	if err != nil {
		fmt.Print(err)
	}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		v := b.Get([]byte(key))
		result = string(v)
		//	fmt.Printf(string(v))
		return nil
	})
	db.Close()
	return result
}
func createImage(w http.ResponseWriter, request *http.Request) {
	user := request.FormValue("user")
	if user != "Kreton" {
		fmt.Print("Incorrect user")
		return
	}
	bucket := "bucket"
	//key := "image"
	//var value = "123test"
	err := request.ParseMultipartForm(100 << 20) // maxMemory 100MB
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

	// bytes are bullshit
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	//fmt.Print(buf)
	Wdb([]byte(bucket), []byte(h.Filename), buf.Bytes())
	//Rdb(bucket, key)
	var a = Rdb(bucket, h.Filename)
	fmt.Print(a)
	tmpfile, err := os.Create("/tmp/" + h.Filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpfile.WriteString(a)
	defer tmpfile.Close()

	_, err = io.Copy(tmpfile, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Print(string(user) + " USER")
	w.WriteHeader(200)

}
