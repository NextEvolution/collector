package main

import (
	"nextevolution/collector/facebookripper"
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
	"nextevolution/collector/config"
)

var cfg config.Config

func main(){

	if len(os.Args) <= 1  || os.Args[1] == ""{
		log.Panic("Please supply a config file path like: ./mock config.json")
	}
	configPath := os.Args[1]

	rawConfig, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Panic(fmt.Sprintf("failed to read config file %s", configPath))
	}

	err = json.Unmarshal(rawConfig, &cfg)
	if err != nil {
		log.Panic(fmt.Sprintf("unable to unmarshal config file %s", configPath))
	}

	ripper := facebookripper.NewFacebookRipper("https://graph.facebook.com/v2.6")

	oauthResp := ripper.GetLongTimeToken( cfg.FbAppId, cfg.FbAppSecret, "some-fb-token" )

	user := ripper.GetUserName("me", oauthResp.AccessToken)

	fmt.Printf("Stuff: %d", user)
}

//func filterSales (sas *collector_dataservice.SellerAlbumScan){
//	fProducts := []collector_dataservice.Product{}
//
//	count := 0
//
//	for _, product := range sas.Products {
//		if len(product.SaleEvents) != 0 {
//			fProducts = append(fProducts, product)
//			count = count + len(product.SaleEvents)
//		}
//	}
//	fmt.Printf("\nFound %d sales\n", count)
//
//	sas.Products = fProducts
//}