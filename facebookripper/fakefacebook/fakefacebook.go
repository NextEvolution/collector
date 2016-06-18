package fakefacebook

import (
	"github.com/onsi/gomega/ghttp"
	"net/http"
)

func NewFakeFacebook() *ghttp.Server{
	server := ghttp.NewServer()

	//setup group endpoint
	server.AppendHandlers(ghttp.CombineHandlers(
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
	return server
}