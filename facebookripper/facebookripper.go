package facebookripper

import (
	"net/http"
	"fmt"
	fb "nextevolution/collector/facebookripper/fbobjects"
	"encoding/json"
	"io/ioutil"
)

func NewFacebookRipper(url string) *FacebookRipper {
	return &FacebookRipper{
		url: url,
	}
}

type FacebookRipper struct {
	url string
}

func (f *FacebookRipper) GetUsersGroups(userId string, token string) []fb.Group {

	url := fmt.Sprintf("%s/%s/groups?access_token=%s",f.url, userId, token)
	resp,err := http.Get(url)

	if err != nil {
		//todo: handle this
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	envelope := fb.GroupEnvelope{}
	json.Unmarshal(body, &envelope)

	return envelope.Data
}
