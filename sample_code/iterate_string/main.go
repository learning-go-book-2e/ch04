package main

import "fmt"

func main() {
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		fmt.Println(sample)
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

}
