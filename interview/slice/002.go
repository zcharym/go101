package main

import (
	"fmt"
)

func SliceRise(s []int) {
	s = append(s, 0)
	for i := range s {
		s[i]++
	}
}
func main() {
	s1 := []int{1, 2}  // s1 {1, 2} cap:3
	s2 := s1           // s2 {1, 2} cap:4 expanded
	s2 = append(s2, 3) // s2 {1, 2, 3}
	SliceRise(s1)      // s1 {1, 2}
	SliceRise(s2)      // s2 {2, 3, 4}
	fmt.Println(s1, s2)
}
