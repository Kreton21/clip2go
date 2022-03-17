package httstuff

import (
	"clip2serv/bolted"
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"time"
)

type application struct {
	auth struct {
		username string
		password string
	}
}

func InitServ() {
	app := new(application)

	app.auth.username = "Kreton"
	app.auth.password = "password"

	if app.auth.username == "" {
		log.Fatal("basic auth username must be provided")
	}
	if app.auth.password == "" {
		log.Fatal("basic auth password must be provided")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/sendimg", app.CreateImage)
	mux.HandleFunc("/sendtxt", app.CreateText)
	mux.HandleFunc("/get", app.test)

	srv := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Printf("starting server on %s", srv.Addr)
	err := srv.ListenAndServeTLS("./localhost.pem", "./localhost-key.pem")
	log.Fatal(err)

	//http.ListenAndServe(":8080", nil)
}

func (app *application) test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the protected handler")
}

func (app *application) CreateText(w http.ResponseWriter, r *http.Request) {
	bucket := "bucket"
	key := "image"
	//var value = "123test"
	err := r.ParseMultipartForm(100 << 20) // maxMemory 100MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Access the photo key - First Approach
	value := r.FormValue("photo")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bolted.Wdb([]byte(bucket), []byte(key), []byte(value))
	var a = bolted.Rdb(bucket, key)

	fmt.Print(a)

	w.WriteHeader(200)
}
func (app *application) CreateImage(w http.ResponseWriter, request *http.Request) {
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
func (app *application) basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(app.auth.username))
			expectedPasswordHash := sha256.Sum256([]byte(app.auth.password))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
