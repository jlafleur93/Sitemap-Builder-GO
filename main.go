package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main(){
	urlFlag := flag.String("url", "https://gophercises.com", "url you want to build a site map for")
	flag.Parse()
	fmt.Println(*urlFlag)
	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
/*
 1. Get Webpage
 2. Parse all the links on the page
 3. Build proper urls with our links
 4. Filter out links
 5. Find all pages via BDS
 6. Print Out XMl
*/ 
