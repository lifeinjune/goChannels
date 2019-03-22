package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	address := []string{"http://google.com", "http://facebook.com", "http://stackoverflow.com", "http://golang.org", "http://amazon.com"} //string slcie contain addresses
	c := make(chan string)                                                                                                                // make the channel type call c
	for _, a := range address {                                                                                                           //loop through each element in slice
		go checkAddress(a, c) //added go child routine and pass c as well
	}
	//fmt.Println(<-c) // main routine will wait for the response from the child routine, and only for fastest child routine complete which not going to wait for other child routine
	for l := range c { //infinite loop
		//rang (channel variable) means that wait for the channel to return value
		go func(address string) { //unnamed function to wrap around delay and no delay start up of child routine (function literal)
			time.Sleep(time.Second)  //delay which pause the current go routine after initial http.get
			checkAddress(address, c) //do not share variables between different routine COPY IT~!
		}(l)
	}
}

func checkAddress(a string, c chan string) { // accept the c channel type with communication type as well which is string

	resp, err := http.Get(a) //get respons and error from address given
	if err != nil {          //if there is error
		fmt.Println("Error!", err) //print out the error statement
		c <- a
		return
	} else {
		fmt.Println(a, "status:", resp.Status) //print out address with it's status
		c <- a
	}
}
