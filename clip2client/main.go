package main

import (
	//	"watchclip"

	"clip2go/httstuff"
	"clip2go/watchclip"
	"sync"
	//	"image/png" //needed to use `png` encoder
	//	"watchclip"
)

// Waits for clipboard to change and sends it to the url
func clipnsend(urlroot string) {
	for {
		value, isImg := watchclip.WatchClip()
		if isImg == true {
			urlroot := urlroot + "/sendimg"
			httstuff.Call(value, urlroot)
		} else {
			urlroot := urlroot + "/sendtxt"
			httstuff.Call(value, urlroot)
		}
	}
}

func main() {
	//Variables to assign with GUI:
	urlroot := "https://localhost:4000"
	//username:= "Kreton"
	//password:= "password"
	var wg sync.WaitGroup //Sync group for go routines
	wg.Add(1)             // Adds a worker to the sync group
	go clipnsend(urlroot)

	wg.Wait() // Wait for all goroutines to finish before exiting
}
