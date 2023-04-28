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
	myLocale       = "en_US"
)

func main() {
	profileClient, err := profile.NewClient(myClientId, myClientSecret, myRegion, myLocale, nil)
	if err != nil {
		log.Fatalln(err)
	}
	res, err := profileClient.CharacterProfileSummary("tichondrius", "sighmir")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
}
