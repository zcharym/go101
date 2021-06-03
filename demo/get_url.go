package main

import (
	"fmt"
	"net/http"
)

func main() {
	const REQ_URL string = "https://example.com"
	res, err := http.Get(REQ_URL)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(res.StatusCode)
}
