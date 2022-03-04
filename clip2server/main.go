package main

import (
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
	http.Handle("/send", createImageHandler)
	http.ListenAndServe(":8080", nil)
}

func Wdb(bucket, key, value string) {
	db, err := bolt.Open("clip2.db", 0600, nil)
	if err != nil {
		fmt.Print(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Put([]byte(key), []byte(value))
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
	bucket := "bucket"
	key := "image"
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
	// Transform multipart to string and insert it into the DB below pls
	Wdb(bucket, key)
	Rdb(bucket, key)
	var a = Rdb(bucket, key)
	fmt.Print(a)
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
