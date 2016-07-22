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

	sas := ripper.GetSoldItems("me", token, "sold", []string{"1601038443544679"}, []string{})

	filterSales(&sas)

	js, _ := json.Marshal(sas)

	fmt.Printf("%s\n\n", js)

	fmt.Println(ripper.CallCount)
}

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