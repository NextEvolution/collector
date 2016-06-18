package facebookripper_test

import (
	. "nextevolution/collector/facebookripper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "./fakefacebook"
	"github.com/onsi/gomega/ghttp"
	"fmt"
)

var _ = Describe("Facebookripper", func() {

	var ripper *FacebookRipper
	var token string
	var userId string
	var server *ghttp.Server

	BeforeEach(func(){
		server = NewFakeFacebook()
		ripper = NewFacebookRipper(server.URL())
		token = "ValidFacebookToken"
		userId = "validUserId"
	})

	It("gets the list of facebook groups", func(){
		groups := ripper.GetUsersGroups(userId, token)
		Expect(len(server.ReceivedRequests())).Should(BeNumerically(">", 0))
		Expect(server.ReceivedRequests()[0].URL.Path).To(Equal("/validUserId/groups"))

		Expect(groups).ToNot(BeNil())
		Expect(groups).ToNot(BeEmpty())
		Expect(groups[0]).ToNot(BeNil())
		fmt.Println(groups)
	})
})
