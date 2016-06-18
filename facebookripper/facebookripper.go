package facebookripper

import (
	"net/http"
	"fmt"
)

func NewFacebookRipper(url string) *FacebookRipper {
	return &FacebookRipper{
		url: url,
	}
}

type FacebookRipper struct {
	url string
}

func (f *FacebookRipper) LookForOrders(userId string, token string){
	url := fmt.Sprintf("%s/%s/groups?access_token=%s",f.url, userId, token)
	fmt.Println("url: " + url)
	http.Get(url)
}
