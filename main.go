package main

import (
	"nextevolution/collector/facebookripper"
	"fmt"
)

func main(){
	ripper := facebookripper.NewFacebookRipper("https://graph.facebook.com/v2.6")
	token := "someToken"

	items := ripper.GetSoldItems("someUserId", token)

	for _, item := range items {
		fmt.Printf("Found Sale: %s, %s, %s, %s\n", item.Photo.Id, item.Photo.Name, item.Comment.From.Name, item.Comment.Message)
	}

	fmt.Println(ripper.CallCount)
}