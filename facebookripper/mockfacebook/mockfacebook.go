package mockfacebook

import (
	"github.com/onsi/gomega/ghttp"
	"net/http"
	"regexp"
)

func NewMockFacebook() *ghttp.Server {
	server := ghttp.NewServer()

	//Groups endpoint
	groupsRegex, _ := regexp.Compile(`\/\w*\/groups`)
	server.RouteToHandler("GET", groupsRegex,
		ghttp.CombineHandlers(
			ghttp.VerifyFormKV("access_token", "ValidFacebookToken"),
			ghttp.RespondWith(http.StatusOK,
				`{
					"data": [
						{
							"name": "Test Group 1",
							"privacy": "CLOSED",
							"id": "testGroupId1"
						}
					],
					"paging": {
						"cursors": {
							"before": "MTA0NDgyNDgxNTU2NDEwNwZDZD",
							"after": "MTYwMTAzODQ0MzU0NDY3OQZDZD"
						}
					}
				}`),
		))

	//Albums Endpoint
	albumsRegex, _ := regexp.Compile(`\/\w*\/albums`)
	server.RouteToHandler("GET", albumsRegex,
		ghttp.CombineHandlers(
			ghttp.VerifyFormKV("access_token", "ValidFacebookToken"),
			ghttp.RespondWith(http.StatusOK,
				`{
					"data": [
						{
							"created_time": "2016-05-19T05:22:50+0000",
							"name": "Test Album",
							"id": "testAlbumId"
						}
					],
					"paging": {
						"cursors": {
							"before": "MTYwMTAzODY4MzU0NDY1NQZDZD",
							"after": "MTYwMTAzODY4MzU0NDY1NQZDZD"
						}
					}
				}`),
		))

	//Photos Endpoint
	photosRegex, _ := regexp.Compile(`\/\w*\/photos`)
	server.RouteToHandler("GET", photosRegex,
		ghttp.CombineHandlers(
			ghttp.VerifyFormKV("access_token", "ValidFacebookToken"),
			ghttp.RespondWith(http.StatusOK,
				`{
				  "data": [
				    {
				      "created_time": "2016-05-19T05:23:53+0000",
				      "name": "Something Something Something",
				      "id": "testPhotoId"
				    }
				  ],
				  "paging": {
				    "cursors": {
				      "before": "MTAxNTQzMDMyODQyODc2MjUZD",
				      "after": "MTAxNTQzMDMyODQ0Mzc2MjUZD"
				    }
				  }
				}`),
		))

	//Comments Endpoint
	commentsRegex, _ := regexp.Compile(`\/\w*\/comments`)
	server.RouteToHandler("GET", commentsRegex,
		ghttp.CombineHandlers(
			ghttp.VerifyFormKV("access_token", "ValidFacebookToken"),
			ghttp.RespondWith(http.StatusOK,
				`{
				  "data": [
				    {
				      "created_time": "2016-05-19T05:43:16+0000",
				      "from": {
				        "name": "Barry Williams",
				        "id": "commenterUserId"
				      },
				      "message": "first comment",
				      "id": "firstCommentId"
				    },
				    {
				      "created_time": "2016-05-19T05:43:16+0000",
				      "from": {
				        "name": "Sally JoBob",
				        "id": "buyerUserId"
				      },
				      "message": "sold",
				      "id": "saleCommentId"
				    }
				  ],
				  "paging": {
				    "cursors": {
				      "before": "WTI5dGJXVnVkRjlqZAFhKemIzSTZANVFl3TVRBME16YzRNelUwTkRFME5Ub3hORFl6TmpNMk5UazIZD",
				      "after": "WTI5dGJXVnVkRjlqZAFhKemIzSTZANVFl3TVRBME16YzRNelUwTkRFME5Ub3hORFl6TmpNMk5UazIZD"
				    }
				  }
				}`),
		))
	return server
}