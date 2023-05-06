package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.cryptosmartnow.io/"

func main() {
	fmt.Println("LCO Web Request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: &T\n", response)
	defer response.Body.Close() //my responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)

		}
		content := string(databytes)
		fmt.Println(content)
	
}

