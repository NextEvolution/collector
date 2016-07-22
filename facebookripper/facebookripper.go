package facebookripper

import (
	"net/http"
	"fmt"
	fb "nextevolution/collector/facebookripper/fbobjects"
	"encoding/json"
	"io/ioutil"
	"regexp"
	. "nextevolution/data-service/types"
	"time"
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
	f.CallCount ++

	resp,err := http.Get(url)

	if err != nil {
		//todo: handle this
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Print("== GET URL ====================================================================\n")
	fmt.Printf("-- Url: GET %s\n", url)
	fmt.Printf("-- Resp Body: %s\n", string(body))
	headers, _ := json.Marshal(resp.Header)
	fmt.Printf("-- Resp Headers: %s\n", string(headers))
	fmt.Printf("-- Resp Code: %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
	fmt.Print("\n")
	return body
}

/* Groups */

func (f *FacebookRipper) handlePagingGroups (urlFormat string, userId string, token string, after string) []fb.Group {

	//go to first url, rip off all data
	body := f.getUrl(urlFormat, userId, token, after)
	envelope := fb.GroupEnvelope{}
	json.Unmarshal(body, &envelope)

	if envelope.Paging.Next == "" {
		return envelope.Data
	}

	// get data for next url
	nextData := f.handlePagingGroups(urlFormat, userId, token, envelope.Paging.Cursors.After)
	return append(envelope.Data, nextData...)
}

func (f *FacebookRipper) GetUsersGroups(userId string, token string) []fb.Group {
	return f.handlePagingGroups("%s/%s/groups?access_token=%s&after=%s", userId, token, "")
}

/* Albums */

func (f *FacebookRipper) handlePagingAlbums (urlFormat string, userId string, token string, after string) []fb.Album {

	//go to first url, rip off all data
	body := f.getUrl(urlFormat, userId, token, after)
	envelope := fb.AlbumEnvelope{}
	json.Unmarshal(body, &envelope)

	if envelope.Paging.Next == "" {
		return envelope.Data
	}

	// get data for next url
	nextData := f.handlePagingAlbums(urlFormat, userId, token, envelope.Paging.Cursors.After)
	return append(envelope.Data, nextData...)
}

func (f *FacebookRipper) GetGroupAlbums(groupId string, token string) []fb.Album {
	return f.handlePagingAlbums("%s/%s/albums?access_token=%s&after=%s&date_format=U", groupId, token, "")
}

/* Photos */

func (f *FacebookRipper) handlePagingPhotos (urlFormat string, userId string, token string, after string) []fb.Photo {

	//go to first url, rip off all data
	body := f.getUrl(urlFormat, userId, token, after)
	envelope := fb.PhotoEnvelope{}
	json.Unmarshal(body, &envelope)

	if envelope.Paging.Next == "" {
		return envelope.Data
	}

	// get data for next url
	nextData := f.handlePagingPhotos(urlFormat, userId, token, envelope.Paging.Cursors.After)
	return append(envelope.Data, nextData...)
}

func (f *FacebookRipper) GetAlbumPhotos(albumId string, token string) []fb.Photo {
	return f.handlePagingPhotos("%s/%s/photos?access_token=%s&after=%s&date_format=U&fields=created_time,name,images,id", albumId, token, "")
}

/* Comments */

func (f *FacebookRipper) handlePagingComments (urlFormat string, userId string, token string, after string) []fb.Comment {

	//go to first url, rip off all data
	body := f.getUrl(urlFormat, userId, token, after)
	envelope := fb.CommentEnvelope{}
	json.Unmarshal(body, &envelope)

	if envelope.Paging.Next == "" {
		return envelope.Data
	}

	// get data for next url
	nextData := f.handlePagingComments(urlFormat, userId, token, envelope.Paging.Cursors.After)
	return append(envelope.Data, nextData...)
}

func (f *FacebookRipper) GetPhotoComments(photoId string, token string) []fb.Comment {
	return f.handlePagingComments("%s/%s/comments?access_token=%s&after=%s&order=chronological&date_format=U", photoId, token, "")
}

/* Main scraping loop */

func (f *FacebookRipper) GetSoldItems(userId string, token string, keyword string, allowedGroups []string, ignoredAlbums []string) SellerAlbumScan {
	sas := SellerAlbumScan{
		Date: int(time.Now().Unix()),
		Products: make([]Product,0,100),
	}

	//iterate over groups
	allGroups := f.GetUsersGroups(userId, token)

	//todo: clean this up. Filters groups to only allowed groups
	var groups []fb.Group
	for _, allowGroup := range allowedGroups {
		for _, group := range allGroups {
			if allowGroup == group.Id {
				groups = append(groups, group)
				break
			}
		}
	}

	for _, group := range groups {

		//iterate over albums
		allAlbums := f.GetGroupAlbums(group.Id, token)

		//filter albums
		var albums []fb.Album

		contains := func(album fb.Album) bool {
			for _, ignoreAlbum := range ignoredAlbums {
				if album.Id == ignoreAlbum {
					return true
				}
			}
			return false
		}

		for _, album := range allAlbums {
			if !contains(album) {
				albums = append(albums, album)
			}
		}

		for _, album := range albums {

			//iterate over photos
			photos := f.GetAlbumPhotos(album.Id, token)
			for _, photo := range photos{

				if len(photo.Images) == 0 {
					fmt.Printf("Unable to read images from picture: %d", photo)
					continue
				}

				//record product
				product := Product {
					Album: album.Name,
					Description: photo.Name,
					Metadata: FbPicture{
						Height: photo.Images[0].Height,
						Width: photo.Images[0].Width,
						ImageUrl: photo.Images[0].Source,
						FbId: photo.Id,
					},
				}

				//iterate over comments
				comments := f.GetPhotoComments(photo.Id, token)

				//fmt.Printf("Comments: %d\n", comments)
				salesEvents := make([]SaleEvent, 0, 0)
				for _, comment := range comments{

					//only record comments if they match the keyword
					if f.Matches(keyword, comment.Message) {

						salesEvents = append(salesEvents, SaleEvent{
							Metadata: FbComment{
								Text: comment.Message,
								FbId: comment.Id,
							},
							Customer: Customer{
								Name: comment.From.Name,
								Metadata: FbUser{
									Name: comment.From.Name,
									FbId: comment.From.Id,
								},
							},
							Date: comment.CreatedTime,
						})
					}
				}

				product.SaleEvents = salesEvents
				sas.Products = append(sas.Products, product)
			}
		}
	}
	return sas
}

/* Helper functions */

func (f *FacebookRipper) Matches(keyword string, message string) bool {
	saleRegex := regexp.MustCompile(`(?i)(\s*|\W)(` + keyword + `)($|\W)`)
	return saleRegex.MatchString(message)
}

//func (f *FacebookRipper) GetLongTimeToken (clientId string, clientSecret string, fbExchangeToken string) {
//	rawResp := f.getUrl("%s/oauth/access_token?grant_type=fb_exchange_token&client_id=%s&client_secret=%s&fb_exchange_token=%s", clientId, clientSecret, fbExchangeToken)
//	fmt.Println(string(rawResp))
//	return ""
//}