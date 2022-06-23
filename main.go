package main

import (
	"errors"
	"fmt"
	"net/http"
)


var errRequestFailed = errors.New("req is failed");

func main(){
	
	var results = make(map[string]string)

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
		result := "OK"
		err := hitURL(url)
		if err != nil{
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results{
		fmt.Println(url, result)
	}
	
}

func hitURL(url string) error {
	fmt.Println("checking : ",url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400{
		return errRequestFailed
	}
	return nil
}