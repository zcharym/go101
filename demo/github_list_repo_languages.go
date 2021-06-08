package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	END_POINT  = "https://api.github.com/repos"
	ACCEPTANCE = "application/vnd.github.v3+json"
)

type LangResult struct {
	C          int `json:"C"`
	Python     int `json:"Python"`
	TypeScript int `json:"TypeScript"`
	JavaScript int `json:"JavaScript"`
	JSON       int `json:"Json"`
	Shell      int `json:"Shell"`
	Go         int `json:"Go"`
	HTML       int `json:"HTML"`
}

func main() {
	var l LangResult
	res := fetch("Zchary-ma", "go101")
	if err := json.NewDecoder(res.Body).Decode(&l); err != nil {
		log.Fatalf("json decode error: %s", err)
	}
	fmt.Printf("%v", l)
}

func fetch(user, repo string) *http.Response {
	var url = END_POINT + "/" + user + "/" + repo + "/languages"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", ACCEPTANCE)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("http get error: %s", err)
	}
	return res
}
