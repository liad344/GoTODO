package Adapters

import "github.com/gen2brain/beeep"

func Notify(title , msg , img string){
	err := beeep.Notify(title , msg, img )
	if err != nil {
		panic(err)
	}
}
