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
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte // not string, for io lib use convenience
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	fmt.Printf("%s", title)
	p, err := loadPage(title)
	t, err := template.ParseFiles("./intro/web_app_view.html")
	if err != nil {
		fmt.Printf("can't find template path")
	} else {
		t.Execute(w, p)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "<h1>Not Found</h1>")
}
func saveHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "<h1>Not Found</h1>")
}
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	// t, err := template.ParseFiles("./intro/web_app_edit.html")
	t, err := template.ParseFiles("./intro/web_app_edit.html")
	if err != nil {
		fmt.Printf("can't find template path")
	} else {
		t.Execute(w, p)
	}

}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/", notFoundHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
