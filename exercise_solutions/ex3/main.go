package main

import "fmt"

func main() {
	var total int
	for i := 0; i < 10; i++ {
		// BUG: the := should be an =
		total := total + i
		fmt.Println(total)
	}
	fmt.Println(total)
}
