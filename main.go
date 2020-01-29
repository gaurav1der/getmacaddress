package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Setting Constant to accept parameters from env

const (
	macAddress = "MACADRESS"
	apiKey     = "API_KEY"
)

// This function will make a call to macaddress api to retrieve OUI vendor information,
// detect virtual machines, possible applications, read the information encoded in the MAC, and get our research's
// results regarding the given MAC address or the OUI.

func main() {
	maddr := os.Getenv(macAddress)
	apik := os.Getenv(apiKey)
	fmt.Println("Entered MAC Address = " + maddr)
	response, err := http.Get("https://api.macaddress.io/v1?apiKey=" + apik + "&output=json&search=" + maddr)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
