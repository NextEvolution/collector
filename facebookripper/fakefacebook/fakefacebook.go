package fakefacebook

import (
	"github.com/onsi/gomega/ghttp"
	"net/http"
)

func NewFakeFacebook() *ghttp.Server {
	server := ghttp.NewServer()

	//Group endpoint
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
						"id": "testAlbum"
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
	return server
}