package main

import (
	"nextevolution/collector/facebookripper"
	"fmt"
)

func main(){
	ripper := facebookripper.NewFacebookRipper("https://graph.facebook.com/v2.6")
	token := "EAACEdEose0cBAHd519jm1Jmc1WKUueHf6TzLufJTtBF9bE2G8ZBOXLapBAPPy3ZCBPT0kBtasijZCMqL3Gq5eK7a7ba65NLqiYxLwVLk86wzZBsJu2nsyZCchraXUHUGSvEtaktJ5wFjTwkp2ntKCN6LmvKMpZBGvwj4tYqZBpZBKAZDZD"

	items := ripper.GetSoldItems("10153843522262625", token, "sold", []string{"1601038443544679"})

	for _, item := range items {
		fmt.Printf("Found Sale: %s, %s, %s, %s, %s\n", item.Photo.Id, item.Photo.Name, item.Comment.Id ,item.Comment.From.Name, item.Comment.Message)
	}

	fmt.Println(ripper.CallCount)
}