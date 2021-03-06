package facebookripper_test

import (
	. "nextevolution/collector/facebookripper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "./mockfacebook"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Facebookripper", func() {

	var ripper *FacebookRipper
	var token string
	var userId string
	var server *ghttp.Server

	BeforeEach(func(){
		server = NewMockFacebook()
		ripper = NewFacebookRipper(server.URL())
		token = "ValidFacebookToken"
		userId = "validUserId"
	})

	Context("Individual Endpoints", func(){
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

	Context("Integration of endpoints", func(){
		It("reports \"sold\" items with customer", func(){
			boughtItems := ripper.GetSoldItems(userId, token, "sold")

			Expect(len(boughtItems)).To(Equal(1))
			Expect(boughtItems[0].Photo.Id).To(Equal("testPhotoId"))
			Expect(boughtItems[0].Comment.From.Id).To(Equal("buyerUserId"))
			Expect(boughtItems[0].Comment.Id).To(Equal("saleCommentId"))
		})
	})

	Context("keyword matcher", func(){
		It("reports \"sold\" with various comment formats", func(){
			//Should match
			Expect(ripper.Matches("sold", "sold")).To(Equal(true))
			Expect(ripper.Matches("sold", "Sold")).To(Equal(true))
			Expect(ripper.Matches("sold", "sOlD")).To(Equal(true))
			Expect(ripper.Matches("sold", "SOLD")).To(Equal(true))
			Expect(ripper.Matches("sold", "     sold")).To(Equal(true))
			Expect(ripper.Matches("sold", "  sold!!!")).To(Equal(true))
			Expect(ripper.Matches("sold", "I want that! SOLD!!!!!")).To(Equal(true))

			//Should not match
			Expect(ripper.Matches("sold", "soldier")).To(Equal(false))
			Expect(ripper.Matches("sold", "solder")).To(Equal(false))
			Expect(ripper.Matches("sold", "sale")).To(Equal(false))

			//Should match other keywords
			//Should match
			Expect(ripper.Matches("happy", "happy")).To(Equal(true))
			Expect(ripper.Matches("happy", "Happy")).To(Equal(true))
			Expect(ripper.Matches("happy", "HaPPy")).To(Equal(true))
			Expect(ripper.Matches("happy", "HAPPY")).To(Equal(true))
			Expect(ripper.Matches("happy", "     happy")).To(Equal(true))
			Expect(ripper.Matches("happy", "  happy!!!")).To(Equal(true))
			Expect(ripper.Matches("happy", "I want that! HAPPY!!!!!")).To(Equal(true))

			//Should not match
			Expect(ripper.Matches("happy", "happyier")).To(Equal(false))
			Expect(ripper.Matches("happy", "happyer")).To(Equal(false))
			Expect(ripper.Matches("happy", "thrilled")).To(Equal(false))

		})
	})
})
