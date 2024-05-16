package main

import (
	// "fmt"
	"math"
)

func GetPrimeNumber(minnumb int, maxnumb int) []int {
	primenumb := make([]int, 0, (maxnumb-minnumb)/2)

	for numb := minnumb; numb <= maxnumb; numb++ {
		flag := 0
		for k := 2; k <= int(math.Sqrt(float64(numb))); k++ {
			if (numb % k) == 0 {
				flag = 1
			}
		}
		if flag == 0 {
			primenumb = append(primenumb, numb)
		}
	} 

	return primenumb
}

// func main() {
// 	fmt.Println(GetPrimeNumber(11, 20))
// }