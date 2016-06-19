package fakefacebook

import (
	"github.com/onsi/gomega/ghttp"
	"net/http"
)

func NewFakeFacebook() *ghttp.Server {
	server := ghttp.NewServer()

	//Groups endpoint
	server.RouteToHandler("GET","/validUserId/groups",ghttp.CombineHandlers(
		ghttp.VerifyRequest("GET", "/validUserId/groups"),
		ghttp.VerifyFormKV("access_token", "ValidFacebookToken"),
		ghttp.RespondWith(http.StatusOK,
			`{
				"data": [
					{
						"name": "Test Group 1",
						"privacy": "CLOSED",
						"id": "testGroupId1"
					},
					{
						"name": "Test Group 2",
						"privacy": "CLOSED",
						"id": "testGroupId2"
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
	server.RouteToHandler("GET", "/testGroupId1/albums",
		ghttp.CombineHandlers(
		ghttp.VerifyRequest("GET", "/testGroupId1/albums"),
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
	server.RouteToHandler("GET", "/testAlbumId/photos",
		ghttp.CombineHandlers(
			ghttp.VerifyRequest("GET", "/testAlbumId/photos"),
			ghttp.VerifyFormKV("access_token", "ValidFacebookToken"),
			ghttp.RespondWith(http.StatusOK,
				`{
				  "data": [
				    {
				      "created_time": "2016-05-19T05:23:53+0000",
				      "name": "Something Something Something",
				      "id": "testPhotoId"
				    },
				    {
				      "created_time": "2016-05-19T05:23:53+0000",
				      "id": "10154303284292625"
				    },
				    {
				      "created_time": "2016-05-19T05:23:53+0000",
				      "id": "10154303284297625"
				    },
				    {
				      "created_time": "2016-05-19T05:23:53+0000",
				      "id": "10154303284302625"
				    },
				    {
				      "created_time": "2016-05-19T05:23:53+0000",
				      "id": "10154303284392625"
				    },
				    {
				      "created_time": "2016-05-19T05:23:53+0000",
				      "id": "10154303284477625"
				    },
				    {
				      "created_time": "2016-05-19T05:23:53+0000",
				      "id": "10154303284397625"
				    },
				    {
				      "created_time": "2016-05-19T05:23:53+0000",
				      "id": "10154303284417625"
				    },
				    {
				      "created_time": "2016-05-19T05:23:53+0000",
				      "id": "10154303284437625"
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
	server.RouteToHandler("GET", "/testPhotoId/comments",
		ghttp.CombineHandlers(
			ghttp.VerifyRequest("GET", "/testPhotoId/comments"),
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
				      "id": "1601043783544145"
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