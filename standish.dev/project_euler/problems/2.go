package problems

import (
	"fmt"
)

// Two returns something
func Two() {
	cur := 0
	next := 1
	sum := 0
	for i := 0; next < 4000000; i++ {
		cur, next = next, cur+next
		if isEven(next) {
			sum += next
		}
	}
	fmt.Println(sum)
}

func isEven(x int) bool {
	return x%2 == 0
}
