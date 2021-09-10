package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"` // omitempty: empty or zero value then not gen json
	Actors []string
}

type SuperHero struct {
	Name           string   `json:"name"`
	Age            int      `json:"age"`
	SecretIdentity string   `json:"secretIdentity"`
	Powers         []string `json:"powers"`
}

func (s SuperHero) String() string {
	return fmt.Sprintf("name:[%s],age:[%d],powers:[%s]", s.Name, s.Age, strings.Join(s.Powers, ","))
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	// marshal: struct to json
	// data, err := json.Marshal(movies)
	// with beautified format
	data, err := json.MarshalIndent(movies, "", "    ")

	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// unmarshal: json to struct
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%s\n", titles)

	// unmarshal from file
	str := parseJSONFile("./assets/json-sample.json")
	fmt.Println("parsed json file:", str)
}

// JSON file parse
// TODO interface support
func parseJSONFile(dir string) (parsedStr string) {
	// store file
	jsonFile, err := os.ReadFile(dir)
	if err != nil {
		fmt.Println(err)
	}
	var heroes []SuperHero
	if err := json.Unmarshal(jsonFile, &heroes); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	for _, hero := range heroes {
		parsedStr += hero.String() + "\n"
	}
	return
}
