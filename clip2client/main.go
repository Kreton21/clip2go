package main

import (
	//	"watchclip"

	"clip2go/httstuff"
	"clip2go/watchclip"
	//	"image/png" //needed to use `png` encoder
	//	"watchclip"
)

func main() {

	for {
		value, isImg := watchclip.WatchClip()
		if isImg == true {
			urlroot := "http://localhost:8080/sendimg"
			httstuff.Call(value, urlroot)
		} else {
			urlroot := "http://localhost:8080/sendtxt"
			httstuff.Call(value, urlroot)
		}
	}
}
