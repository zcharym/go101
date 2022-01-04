package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"sync"
)

/**
 * Simple MapReduce Job
 *
 */

type PlayerInfo struct {
	Name        string `json:"Name"`
	Nationality string `json:"Nationality"`
	Rating      int    `json:"Rating"`
	Age         int    `json:"Age"`
}

func main() {

	// open CSV file and read into lines
	path := path.Join(os.Getenv("DOWNLOAD_DIR"), "./FullData.csv")
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// initial mapper list
	lists := make(chan []PlayerInfo)

	// after reducer proceeds
	output := make(chan []PlayerInfo)

	var wg sync.WaitGroup
	wg.Add(len(lines) - 1)

	for idx, line := range lines {
		if idx == 0 {
			continue
		}
		go func(item []string) {
			defer wg.Done()
			lists <- Map(item)
		}(line)
	}

	go Reduce(lists, output)

	wg.Wait()
	close(lists)
	fmt.Println(<-output)

}

// Map Mapper implementation
func Map(item []string) []PlayerInfo {
	list := []PlayerInfo{}
	rating, _ := strconv.Atoi(item[9])
	age, _ := strconv.Atoi(item[14])
	list = append(list, PlayerInfo{
		Name:        item[0],
		Nationality: item[1],
		Rating:      rating,
		Age:         age,
	})
	return list
}

// Reduce Reducer impelmentation
func Reduce(input, output chan []PlayerInfo) {
	result := []PlayerInfo{}
	for list := range input {
		for _, v := range list {
			if v.Age >= 25 {
				result = append(result, v)
			}
		}
	}
	output <- result
}
