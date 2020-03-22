package problems

import "fmt"

// One returns something
func One() {
	sum := 0
	for i := range [1000]int{} {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
