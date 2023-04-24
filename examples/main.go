package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/guicaulada/wowapi/pkg/profile"
)

const (
	myClientId     = "myClientId"
	myClientSecret = "myClientSecret"
	myRegion       = "us"
)

func main() {
	profileClient, err := profile.NewClient(myClientId, myClientSecret, myRegion, nil)
	if err != nil {
		log.Fatalln(err)
	}
	res, err := profileClient.CharacterProfileSummary("tichondrius", "sighmir", "profile-us", "en_US")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
}
