package main

import "fmt"

func main() {
	// //simple iterations
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println("for loops", i)
	// }
	//iteration over collection
	// numbers := []int{1, 2, 3, 4, 5}
	// for index, value := range numbers {
	// 	fmt.Printf(" index: %d , value: %d", index, value)
	// }

	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue // skips the rest of the loop when i is 3
		}
		if i == 5 {
			break // exits the loop when i is 5
		}
		fmt.Println("for loops", i)
	}
}
