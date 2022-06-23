package main

import (
	"errors"
	"fmt"
	"net/http"
)

// func main(){
// 	c := make(chan string)
// 	people := [2]string{"jiwan","test"}

// 	for _,person := range people{
// 		go testCount(person,c)
// 	}

// 	for i:=0;i<len(people);i++{

// 		fmt.Println(<-c)
// 	}

// 	fmt.Println("waiting");

// }

// func testCount(word string, c chan string){
// 	time.Sleep(time.Second * 5)
// 	c <- word + " is finished"
// }



var errRequestFailed = errors.New("req is failed");

type result struct{
	url string
	flag string
}

func main(){

	
	results := make(map[string]string)

	c := make(chan result)
	
	

	urls := []string{
		"https://www.airbnb.com/",
"https://www.google.com/",
"https://www.amazon.com/",
"https://www.reddit.com/",
"https://www.google.com/",
"https://soundcloud.com/",
"https://www.facebook.com/",
"https://www.instagram.com/",
"https://academy.nomadcoders.co/",
	}

	// result := map[string]string

	for _,url := range urls {
		
		go hitURL(url,c)
		
	}

	for i := 0;i<len(urls);i++{
		reqResult := <-c
		results[reqResult.url] = reqResult.flag
	}
	
	for url,status := range results{
		fmt.Println(url, " is ",status)
	}
	
}

func hitURL(url string, c chan<- result) {
	
	resp, err := http.Get(url)
	status := "OK"
	
	if err != nil || resp.StatusCode >= 400{
		status = "FAILED"
	}

	c <- result{url:url, flag:status}
	
}