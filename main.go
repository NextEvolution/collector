package main

import (
	"nextevolution/collector/facebookripper"
	"fmt"
	"encoding/json"
	"nextevolution/common_types/collector_dataservice"
)

func main(){
	ripper := facebookripper.NewFacebookRipper("https://graph.facebook.com/v2.6")
	token := "EAAIhEU5mG54BAFS4vU8Us09v0WD1QuNe5gj8iZA5G16KKlxLJZCS97UzJZCH5rM5LDqqPjlOHx1D655VoDMdUBeYmyxFYGKL039EqebmPAZBZAQ07pofNZAZB315kyHWIx2TdTGKD60TZBhji06ZCmVmLqWqummHLqNIZD"

	sas := ripper.GetSoldItems("me", token, "sold", []string{"1601038443544679"}, []string{"1617591905222666"})

	//sas := ripper.GetSoldItems("me", token, "sold", []string{"1044824815564107"}, []string{})

	filterSales(&sas)

	js, _ := json.Marshal(sas)

	fmt.Printf("%s\n\n", js)

	//for _, item := range items {
	//	fmt.Printf("Found Sale: %s, %s, %s, %s, %s\n", item.Photo.Id, item.Photo.Name, item.Comment.Id ,item.Comment.From.Name, item.Comment.Message)
	//}

	fmt.Println(ripper.CallCount)

	//ripper.GetLongTimeToken("599308166896542", "client secret", "EAAIhEU5mG54BAG0RIgNuplobqy8FoD6Kc66cHMnFMHvHyeI2ps8CwQ5zLb14CL9UAPbNIMZBdVpJl7Rx3aT7r85iMdsOXQyW5OqaZBDQZCC54mhXwok0RFehKgMvys3BYARpN6pWL8ZBFSq4AXRJKiXBN38qlNULdJDQBP0exAZDZD")
}

//func RipSalesToNats (userId string, token string, keyword string, allowedGroups []string){
//	ripper := facebookripper.NewFacebookRipper("https://graph.facebook.com/v2.6")
//	items := ripper.GetSoldItems("10153843522262625", token, "sold", []string{"1601038443544679"}, []string{"1601038683544655"})
//
//}

func filterSales (sas *collector_dataservice.SellerAlbumScan){
	fProducts := []collector_dataservice.Product{}

	count := 0

	for _, product := range sas.Products {
		if len(product.SaleEvents) != 0 {
			fProducts = append(fProducts, product)
			count = count + len(product.SaleEvents)
		}
	}
	fmt.Printf("\nFound %d sales\n", count)

	sas.Products = fProducts
}