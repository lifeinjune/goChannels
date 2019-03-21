package main

import (
	"fmt"
	"net/http"
)

func main() {
	address := []string{"http://google.com", "http://facebook.com", "http://stackoverflow.com", "http://golang.org", "http://amazon.com"} //string slcie contain addresses
	for _, a := range address {                                                                                                           //loop through each element in slice
		checkAddress(a)
	}
}

func checkAddress(a string) {
	resp, err := http.Get(a) //get respons and error from address given
	if err != nil {          //if there is error
		fmt.Println("Error!", err) //print out the error statement
		return
	} else { //otherwise
		fmt.Println(a, "is", resp.Status) //print out address with it's status
	}
}
