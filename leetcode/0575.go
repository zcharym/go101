package leetcode

func distributeCandies(candyType []int) int {
	candyLength := len(candyType)
	m := make(map[int]int)
	for _, v := range candyType {
		m[v]++
	}

	count := len(m)
	if candyLength/2 >= count {
		return count
	} else {
		return candyLength / 2
	}
}
