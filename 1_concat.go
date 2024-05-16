package main

import (
	"fmt"
)

func GetComputerString(number int) string {
	last_numb := number % 10
	if last_numb == 1 {
		return fmt.Sprintf("%d компьютер", number)
	} else if last_numb > 1 && last_numb < 5 {
		return fmt.Sprintf("%d компьютера", number)
	} else {
		return fmt.Sprintf("%d компьютеров", number)
	}
}

// func main() {
// 	fmt.Println(GetComputerString(25))
// 	fmt.Println(GetComputerString(41))
// 	fmt.Println(GetComputerString(1048))
// }