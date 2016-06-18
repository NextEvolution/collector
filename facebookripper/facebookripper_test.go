package facebookripper_test

import (
	. "nextevolution/collector/facebookripper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"net/http"
)

var _ = Describe("Facebookripper", func() {
	var server *ghttp.Server
	var ripper *FacebookRipper
	var token string
	var userId string

	BeforeEach(func(){
		server = ghttp.NewServer()
		ripper = NewFacebookRipper(server.URL())
		token = "ValidFacebookToken"
		userId = "validUserId"

		//setup server
		server.AppendHandlers(ghttp.CombineHandlers(
			ghttp.VerifyRequest("GET", "/validUserId/groups"),
			ghttp.VerifyFormKV("access_token", "ValidFacebookToken"),
			ghttp.RespondWith(http.StatusOK, "something"),
		))
	})

	It("calls the group url", func(){
		ripper.LookForOrders(userId, token)
		Expect(len(server.ReceivedRequests())).To(Equal(1))
	})
})
