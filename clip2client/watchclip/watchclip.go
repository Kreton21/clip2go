package watchclip

import (
	//	"watchclip"

	"context"

	//	"image/png" //needed to use `png` encoder

	"golang.design/x/clipboard"
)

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
			isImg := true
			return a, isImg
		}
	case b := <-changedTxt:
		{
			isImg := true
			return b, isImg
		}
	}
}
