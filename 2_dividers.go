package main

// import (
// 	"fmt"
// )

func TotalDividers(arr []int) []int {
	dividers := make(map[int]int)
	total_dividers := make([]int, 0, 8)

	total_dividers = append(total_dividers, 1)

	for i := 0; i < len(arr); i++ {
		elem := arr[i]
		border := elem / 2 + 1

		for k := 2; k <= int(border); k++ {
			if (elem % k) == 0 {
				value := dividers[k]
				value++
				dividers[k] = value
			}
		}
	} 
	
	for numb, count := range dividers {
		if count == 3 {
			total_dividers = append(total_dividers, numb)
		}
	}
	
	return total_dividers 
}

// func main() {
// 	fmt.Println(TotalDividers([]int{42, 12, 18}))
// }