package httstuff

import (
	"bytes"
	"log"
	"mime/multipart"
	"net/http"
	"time"
)

func Call(filename []byte, urlPath string) error {

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	photo, err := writer.CreateFormField("photo")
	if err != nil {
		return err
	}
	photo.Write([]byte(filename))

	writer.Close()
	req, err := http.NewRequest("Post", urlPath, bytes.NewReader(body.Bytes()))
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
