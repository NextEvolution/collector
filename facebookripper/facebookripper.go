package facebookripper

import (
	"net/http"
	"fmt"
	fb "nextevolution/collector/facebookripper/fbobjects"
	"encoding/json"
	"io/ioutil"
	"regexp"
)

func NewFacebookRipper(url string) *FacebookRipper {
	return &FacebookRipper{
		url: url,
	}
}

type FacebookRipper struct {
	url       string
	CallCount int
}

func (f *FacebookRipper) GetUsersGroups(userId string, token string) []fb.Group {

	url := fmt.Sprintf("%s/%s/groups?access_token=%s",f.url, userId, token)
	f.CallCount ++

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
	f.CallCount ++

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

	url := fmt.Sprintf("%s/%s/photos?&access_token=%s",f.url, albumId, token)
	f.CallCount ++

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
	f.CallCount ++

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


func (f *FacebookRipper) Matches(keyword string, message string) bool {
	saleRegex := regexp.MustCompile(`(?i)(\s*|\W)(` + keyword + `)($|\W)`)
	return saleRegex.MatchString(message)
}

func (f *FacebookRipper) GetSoldItems(userId string, token string, keyword string) []*Sale {
	var sales []*Sale

	//iterate over groups
	groups := f.GetUsersGroups(userId, token)
	for _, group := range groups {
		fmt.Printf("Group: %s, %s\n", group.Id, group.Name)

		//iterate over albums
		albums := f.GetGroupAlbums(group.Id, token)
		for _, album := range albums {
			fmt.Printf("Album: %s, %s\n", album.Id, album.Name)

			//iterate over photos
			photos := f.GetAlbumPhotos(album.Id, token)
			for _, photo := range photos{
				fmt.Printf("Photo: %s, %s, %s\n", photo.Id, photo.Name,photo.CreatedTime)

				//iterate over comments
				comments := f.GetPhotoComments(photo.Id, token)
				for _, comment := range comments{
					fmt.Printf("Comment: %s, %s\n", comment.Id, comment.From.Name)

					//found a sale
					if f.Matches(keyword, comment.Message) {
						fmt.Println("found a sale!")
						sale := &Sale{
							Photo: photo,
							Comment: comment,
						}
						sales = append(sales, sale)
						break
					}
				}
			}
		}
	}
	return sales
}