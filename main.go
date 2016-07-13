package main

import (
	"nextevolution/collector/facebookripper"
	"fmt"
)

func main(){
	ripper := facebookripper.NewFacebookRipper("https://graph.facebook.com/v2.6")
	token := "EAACEdEose0cBAJpMXknSnYJn4ukZCybZBZCzwWw5UQ7jrEbdgsNnKqCQFqQ6bnNOJwEQmWZBJaXyiqsFCZB6ZBBi2GqkE3uZAW6GR5AlFqZCobI96NZATJa1mYsaxfICqnrIICbzkgrIngj6LBMLPs3e8ZCZBuoJ426HjBsF1ZAw3m7NDgZDZD"

	items := ripper.GetSoldItems("10153843522262625", token, "sold")

	for _, item := range items {
		fmt.Printf("Found Sale: %s, %s, %s, %s, %s\n", item.Photo.Id, item.Photo.Name, item.Comment.Id ,item.Comment.From.Name, item.Comment.Message)
	}

	fmt.Println(ripper.CallCount)
}