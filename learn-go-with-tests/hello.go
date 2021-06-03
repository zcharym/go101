package hello

import "fmt"

func main() {
	user := "Alex"
	lang := "Chinese"
	fmt.Println(greeting(user, lang))
}

func greeting(name, language string) string {
	if name == "" {
		name = "Tom"
	}
	if language == "" {
		language = "Chinese"
	}
	return "hello, " + name + " in " + language
}
