package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html"
)

const (
	baseURL = "https://redis.io"
)

type RedisDoc struct {
}

func main() {
	// Get base node
	basenode, err := FetchNode(baseURL + "/commands")
	if err != nil {
		log.Printf("%v", err)
	}

	// Find all the links
	var (
		links   []string
		crawler func(*html.Node)
		wg      sync.WaitGroup
	)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			flag, link := expectedLink(node)
			if flag {
				links = append(links, link)
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(basenode)

	result := make(map[string]*html.Node, len(links))
	for _, link := range links {
		url := baseURL + link
		link := link
		wg.Add(1)
		log.Printf("link: %s", link)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * 2)
			node, err := FetchNode(url)
			if err != nil {
				log.Printf("%v", err)
			}
			result[link] = node
		}()
	}

	wg.Wait()
	log.Printf("%+v", result)
}

func FetchNode(url string) (*html.Node, error) {
	client := &http.Client{}

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	node, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func expectedLink(node *html.Node) (bool, string) {
	var (
		flag bool
		link string
	)

	for _, attr := range node.Attr {
		if attr.Key == "class" && attr.Val == "absolute inset-0 outline-0" {
			flag = true
		}

		if attr.Key == "href" {
			link = attr.Val
		}
	}
	return flag, link
}
