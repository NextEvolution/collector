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

func (f *FacebookRipper) getUrl (urlFormat string, userId string, token string, after string) []byte {
	url := fmt.Sprintf(urlFormat, f.url, userId, token, after)
	fmt.Printf("getUrl: %s", url)
	f.CallCount ++

	resp,err := http.Get(url)

	if err != nil {
		//todo: handle this
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func (f *FacebookRipper) handlePaging (urlFormat string, userId string, token string, after string) []interface{} {
	fmt.Println("")

	//go to first url, rip off all data
	body := f.getUrl(urlFormat, userId, token, after)

	//fmt.Println(string(body))

	envelope := fb.Envelope{}
	json.Unmarshal(body, &envelope)

	if envelope.Paging.Next == "" {
		fmt.Printf("Length of data: %d\n", len(envelope.Data))
		return envelope.Data
	}

	// get data for next url
	nextData := f.handlePaging(urlFormat, userId, token, envelope.Paging.Cursors.After)

	return append(envelope.Data, nextData...)
}

func (f *FacebookRipper) getData(urlFormat string, userId string, token string) []interface{} {
	return f.handlePaging(urlFormat, userId, token, "")
}

func (f *FacebookRipper) GetUsersGroups(userId string, token string) []fb.Group {

	rawData := f.getData("%s/%s/groups?access_token=%s&after=%s", userId, token)

	var groups []fb.Group
	groups = make([]fb.Group, len(rawData), len(rawData))

	for i, data := range rawData {
		js, err := json.Marshal(data)
		if err!= nil {
			panic ("I got an error trying to marshall")
		}
		gr := fb.Group{}
		json.Unmarshal(js, &gr)
		if err!= nil {
			panic ("I got an error trying to unmarshall")
		}
		groups[i] = gr
	}
	return groups
}

func (f *FacebookRipper) GetGroupAlbums(groupId string, token string) []fb.Album {
	rawData := f.getData("%s/%s/albums?access_token=%s&after=%s", groupId, token)

	var albums []fb.Album
	albums = make([]fb.Album, len(rawData), len(rawData))

	for i, data := range rawData {
		js, err := json.Marshal(data)
		if err!= nil {
			panic ("I got an error trying to marshall")
		}
		al := fb.Album{}
		json.Unmarshal(js, &al)
		if err!= nil {
			panic ("I got an error trying to unmarshall")
		}
		albums[i] = al
	}
	return albums
}

func (f *FacebookRipper) GetAlbumPhotos(albumId string, token string) []fb.Photo {
	rawData := f.getData("%s/%s/photos?access_token=%s&after=%s", albumId, token)

	var photos []fb.Photo
	photos = make([]fb.Photo, len(rawData), len(rawData))

	for i, data := range rawData {
		js, err := json.Marshal(data)
		if err!= nil {
			panic ("I got an error trying to marshall")
		}
		ph := fb.Photo{}
		json.Unmarshal(js, &ph)
		if err!= nil {
			panic ("I got an error trying to unmarshall")
		}
		photos[i] = ph
	}
	return photos
}

func (f *FacebookRipper) GetPhotoComments(photoId string, token string) []fb.Comment {
	rawData := f.getData("%s/%s/comments?access_token=%s&after=%s", photoId, token)

	var comments []fb.Comment
	comments = make([]fb.Comment, len(rawData), len(rawData))

	for i, data := range rawData {
		js, err := json.Marshal(data)
		if err!= nil {
			panic ("I got an error trying to marshall")
		}
		com := fb.Comment{}
		json.Unmarshal(js, &com)
		if err!= nil {
			panic ("I got an error trying to unmarshall")
		}
		comments[i] = com
	}
	return comments
}


func (f *FacebookRipper) Matches(keyword string, message string) bool {
	saleRegex := regexp.MustCompile(`(?i)(\s*|\W)(` + keyword + `)($|\W)`)
	return saleRegex.MatchString(message)
}

func (f *FacebookRipper) GetSoldItems(userId string, token string, keyword string) []*Sale {
	var sales []*Sale

	//iterate over groups
	groups := f.GetUsersGroups(userId, token)
	for _, g := range groups {
		fmt.Printf("Group: %s, %s\n", g.Id, g.Name)
	}
	for _, group := range groups {
		fmt.Printf("Group: %s, %s\n", group.Id, group.Name)



		//iterate over albums
		albums := f.GetGroupAlbums(group.Id, token)
		for _, album := range albums {
			//if album.Id != "1601038683544655" {
			//	continue
			//}
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