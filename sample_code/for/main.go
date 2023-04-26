package main

import "fmt"

func main() {
	completeFor()
}

func completeFor() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func conditionOnlyFor() {
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}
}
