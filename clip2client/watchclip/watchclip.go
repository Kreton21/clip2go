package watchclip

import (
	//	"watchclip"

	"context"
	"fmt"

	//	"image/png" //needed to use `png` encoder

	"golang.design/x/clipboard"
)

func _init() {
	fmt.Println("Clip Watcher 0.0.2")
}
func WatchClip() ([]byte, bool) {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	changedImg := clipboard.Watch(context.Background(), clipboard.FmtImage)
	changedTxt := clipboard.Watch(context.Background(), clipboard.FmtText)
	select {
	case a := <-changedImg:
		{
			fmt.Println("Recieved Image")
			isImg := true
			return a, isImg
		}
	case b := <-changedTxt:
		{
			fmt.Println("Recieved Text")
			isImg := true
			return b, isImg
		}
	}
}
