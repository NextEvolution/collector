package main

import (
	"encoding/json"
	"github.com/nats-io/nats"
	"fmt"
	"nextevolution/collector/facebookripper"
	"nextevolution/collector/types"
)

var nc *nats.Conn
var ripper *facebookripper.FacebookRipper
func main(){
	killCh := make(chan bool, 1)

	var err error
	nc, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(fmt.Sprintf("Cannot connect to NATS: %s", err))
	}

	ripper = facebookripper.NewFacebookRipper("https://graph.facebook.com/v2.6")

	Serve()
	<- killCh
}

func Serve(){
	fmt.Println("listening ...")
	nc.Subscribe("collector", HandleRequest)
}

func HandleRequest(m *nats.Msg){
	fmt.Printf("Received a message: %s\n", string(m.Data))

	var request types.Request
	err := json.Unmarshal(m.Data, &request)
	if err != nil {
		fmt.Printf("Error, unable to unmarshal request: %s", string(m.Data))
	}

	sas := ripper.GetSoldItems("me", request.FbToken, request.Keywords[0], request.Groups, request.IgnoreAlbums)

	js, _ := json.Marshal(sas)
	fmt.Printf("sending: %s", string(js))
	nc.Publish("data-service", js)
}