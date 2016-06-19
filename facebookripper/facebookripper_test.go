package facebookripper_test

import (
	. "nextevolution/collector/facebookripper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "./fakefacebook"
	"github.com/onsi/gomega/ghttp"
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
		Expect(groups[0].Id).To(Equal("testGroupId1"))
	})

	It("gets the albums of a group", func(){
		groupId := "testGroupId1"
		albums := ripper.GetGroupAlbums(groupId, token)

		Expect(len(server.ReceivedRequests())).Should(BeNumerically(">", 0))
		Expect(server.ReceivedRequests()[0].URL.Path).To(Equal("/testGroupId1/albums"))

		Expect(albums).ToNot(BeNil())
		Expect(albums).ToNot(BeEmpty())
		Expect(albums[0]).ToNot(BeNil())
		Expect(albums[0].Id).To(Equal("testAlbumId"))
	})

	It("gets a list of photos in an album", func(){
		albumId := "testAlbumId"
		photos := ripper.GetAlbumPhotos(albumId, token)

		Expect(len(server.ReceivedRequests())).Should(BeNumerically(">", 0))
		Expect(server.ReceivedRequests()[0].URL.Path).To(Equal("/testAlbumId/photos"))

		Expect(photos).ToNot(BeNil())
		Expect(photos).ToNot(BeEmpty())
		Expect(photos[0]).ToNot(BeNil())
		Expect(photos[0].Id).To(Equal("testPhotoId"))
	})

	It("gets the comments on a photo", func (){
		photoId := "testPhotoId"
		comments := ripper.GetPhotoComments(photoId, token)

		Expect(len(server.ReceivedRequests())).Should(BeNumerically(">", 0))
		Expect(server.ReceivedRequests()[0].URL.Path).To(Equal("/testPhotoId/comments"))

		Expect(comments).ToNot(BeNil())
		Expect(comments).ToNot(BeEmpty())
		Expect(comments[0]).ToNot(BeNil())
		Expect(comments[0].Id).To(Equal("firstCommentId"))
		Expect(comments[0].From.Id).To(Equal("commenterUserId"))
		Expect(comments[0].Message).To(Equal("first comment"))
	})
})
