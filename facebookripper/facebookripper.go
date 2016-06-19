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

func (f *FacebookRipper) GetGroupAlbums(groupId string, token string) []fb.Album {

	url := fmt.Sprintf("%s/%s/albums?access_token=%s",f.url, groupId, token)
	resp,err := http.Get(url)

	if err != nil {
		//todo: handle this
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	envelope := fb.AlbumEnvelope{}
	json.Unmarshal(body, &envelope)

	return envelope.Data
}

func (f *FacebookRipper) GetAlbumPhotos(albumId string, token string) []fb.Photo {

	url := fmt.Sprintf("%s/%s/photos?access_token=%s",f.url, albumId, token)
	resp,err := http.Get(url)

	if err != nil {
		//todo: handle this
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	envelope := fb.PhotoEnvelope{}
	json.Unmarshal(body, &envelope)

	return envelope.Data
}

func (f *FacebookRipper) GetPhotoComments(photoId string, token string) []fb.Comment {

	url := fmt.Sprintf("%s/%s/comments?access_token=%s",f.url, photoId, token)
	resp,err := http.Get(url)

	if err != nil {
		//todo: handle this
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	envelope := fb.CommentEnvelope{}
	json.Unmarshal(body, &envelope)

	return envelope.Data
	return nil
}