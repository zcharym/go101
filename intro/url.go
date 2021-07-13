package main

import (
	"fmt"
	"log"
	url2 "net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln(">>> input url first ")
	}

	urlStr := os.Args[2]
	url, err := url2.Parse(urlStr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("url info: %#v", url)

}
