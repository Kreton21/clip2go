package main

import (
	//	"watchclip"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	//	"image/png" //needed to use `png` encoder
	"os"
	"strconv"

	//	"watchclip"
	"golang.design/x/clipboard"
)

func call(filename, urlPath, method string) error {

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("photo", filename)
	if err != nil {
		return err
	}
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		return err
	}
	user, err := writer.CreateFormField("user")
	user.Write([]byte("Kretonn"))
	if err != nil {
		return err
	}
	writer.Close()
	req, err := http.NewRequest(method, urlPath, bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", rsp.StatusCode)
	}
	return nil
}
func main() {
	urlroot := "http://localhost:8080"
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	i := 0

	var a string

	for {
		changed := clipboard.Watch(context.Background(), clipboard.FmtImage)
		a = string(<-changed)
		filename := strconv.Itoa(i) + " image.png"
		img, err := os.Create(filename)
		fmt.Println("created")

		img.WriteString(a)
		fmt.Println("written " + strconv.Itoa(i))

		call(filename, urlroot+"/send", "POST")
		img.Close()
		os.Remove(filename)

		if err != nil {
			fmt.Println(err)
		}
		if i == 4 {
			i = 0
		} else {
			i += 1
		}
		//		fmt.Println(string(<-changed))
	}
}
