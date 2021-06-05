/**
- Creating a data structure with load and save methods
- Using the net/http package to build web applications
- Using the html/template package to process HTML templates
- Using the regexp package to validate user input
- Using closures
*/

package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte // not string, for io lib use convenience
}

func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	err := p1.Save()
	if err != nil {
		fmt.Println(err.Error())
	}
	p2, _ := LoadPage("TestPage")
	fmt.Println(string(p2.Body))
}
