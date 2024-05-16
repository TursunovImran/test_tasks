package main

import (
	"fmt"
)

func PrintMultiplicationTable(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			product := i * j
			fmt.Printf("%d\t", product)
		}
		fmt.Println()
	}
}

// func main() {
// 	PrintMultiplicationTable(5)
// }